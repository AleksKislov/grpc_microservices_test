package main

import (
    "context"
    "log"
    "net"
    "sync"
    "time"
		"fmt"

    pb "github.com/AleksKislov/grpc_microservices_test/proto/order"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type orderService struct {
    pb.UnimplementedOrderServiceServer
    mu     sync.RWMutex
    orders map[string]*pb.Order
}

func newOrderService() *orderService {
    return &orderService{
        orders: make(map[string]*pb.Order),
    }
}

func (s *orderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Simple ID generation
    id := fmt.Sprintf("order_%d", len(s.orders)+1)

    order := &pb.Order{
        Id:        id,
        UserId:    req.UserId,
        Items:     req.Items,
        Status:    "pending",
        CreatedAt: time.Now().Format(time.RFC3339),
    }

    // Calculate total amount
    var totalAmount float32
    for _, item := range req.Items {
        totalAmount += item.Price * float32(item.Quantity)
    }
    order.TotalAmount = totalAmount

    s.orders[id] = order

    return &pb.OrderResponse{Order: order}, nil
}

func (s *orderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    order, exists := s.orders[req.Id]
    if !exists {
        return nil, status.Errorf(codes.NotFound, "order not found")
    }

    return &pb.OrderResponse{Order: order}, nil
}

func (s *orderService) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    var userOrders []*pb.Order
    for _, order := range s.orders {
        if order.UserId == req.UserId {
            userOrders = append(userOrders, order)
        }
    }

    return &pb.ListOrdersResponse{
        Orders: userOrders,
        Total:  int32(len(userOrders)),
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    server := grpc.NewServer()
    pb.RegisterOrderServiceServer(server, newOrderService())

    log.Println("Starting order service on :50052")
    if err := server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

