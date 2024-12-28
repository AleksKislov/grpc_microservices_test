package main

import (
    "context"
    "fmt"
    "log"
		"os"
    "net"
    "sync"
    "time"

    paymentPb "github.com/AleksKislov/grpc_microservices_test/proto/payment"
    orderPb "github.com/AleksKislov/grpc_microservices_test/proto/order"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type paymentService struct {
    paymentPb.UnimplementedPaymentServiceServer
    mu        sync.RWMutex
    payments  map[string]*paymentPb.Payment
    orderClient orderPb.OrderServiceClient
}

func newPaymentService(orderClient orderPb.OrderServiceClient) *paymentService {
    return &paymentService{
        payments: make(map[string]*paymentPb.Payment),
        orderClient: orderClient,
    }
}

func (s *paymentService) ProcessPayment(ctx context.Context, req *paymentPb.ProcessPaymentRequest) (*paymentPb.PaymentResponse, error) {
    orderReq := &orderPb.GetOrderRequest{
        Id: req.OrderId,
    }
    
    _, err := s.orderClient.GetOrder(ctx, orderReq)
    if err != nil {
        return nil, status.Errorf(codes.InvalidArgument, "order not found: %v", err)
    }

    s.mu.Lock()
    defer s.mu.Unlock()

    paymentId := fmt.Sprintf("payment_%d", len(s.payments)+1)

    payment := &paymentPb.Payment{
        Id:            paymentId,
        OrderId:       req.OrderId,
        UserId:        req.UserId,
        Amount:        req.Amount,
        Status:        "processing",
        PaymentMethod: req.PaymentMethod,
        CreatedAt:     time.Now().Format(time.RFC3339),
    }

    payment.Status = "completed"

    updateOrderReq := &orderPb.UpdateOrderRequest{
        Id:     req.OrderId,
        Status: "confirmed",
    }

    _, err = s.orderClient.UpdateOrder(ctx, updateOrderReq)
    if err != nil {
        payment.Status = "failed"
        s.payments[paymentId] = payment
        return nil, status.Errorf(codes.Internal, "failed to update order status: %v", err)
    }

    s.payments[paymentId] = payment

    return &paymentPb.PaymentResponse{Payment: payment}, nil
}

func (s *paymentService) GetPaymentStatus(ctx context.Context, req *paymentPb.GetPaymentStatusRequest) (*paymentPb.PaymentResponse, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    payment, exists := s.payments[req.PaymentId]
    if !exists {
        return nil, status.Errorf(codes.NotFound, "payment not found")
    }

    return &paymentPb.PaymentResponse{Payment: payment}, nil
}

func main() {
    orderServiceAddr := os.Getenv("ORDER_SERVICE_ADDR")
    orderConn, err := grpc.Dial(orderServiceAddr, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to connect to order service: %v", err)
    }
    defer orderConn.Close()

    orderClient := orderPb.NewOrderServiceClient(orderConn)

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    paymentPb.RegisterPaymentServiceServer(grpcServer, newPaymentService(orderClient))

    log.Println("Starting payment service on :50053")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
