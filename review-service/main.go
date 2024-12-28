package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "sync"
    "time"

    reviewPb "github.com/AleksKislov/grpc_microservices_test/proto/review"
    userPb "github.com/AleksKislov/grpc_microservices_test/proto/user"
    orderPb "github.com/AleksKislov/grpc_microservices_test/proto/order"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type reviewService struct {
    reviewPb.UnimplementedReviewServiceServer
    mu          sync.RWMutex
    reviews     map[string]*reviewPb.Review
    userClient  userPb.UserServiceClient
    orderClient orderPb.OrderServiceClient
}

func newReviewService(userClient userPb.UserServiceClient, orderClient orderPb.OrderServiceClient) *reviewService {
    return &reviewService{
        reviews:     make(map[string]*reviewPb.Review),
        userClient:  userClient,
        orderClient: orderClient,
    }
}

func (s *reviewService) CreateReview(ctx context.Context, req *reviewPb.CreateReviewRequest) (*reviewPb.ReviewResponse, error) {
    userReq := &userPb.GetUserRequest{
        Id: req.UserId,
    }
    
    _, err := s.userClient.GetUser(ctx, userReq)
    if err != nil {
        return nil, status.Errorf(codes.InvalidArgument, "user not found: %v", err)
    }

    orderReq := &orderPb.GetOrderRequest{
        Id: req.OrderId,
    }
    
    orderResp, err := s.orderClient.GetOrder(ctx, orderReq)
    if err != nil {
        return nil, status.Errorf(codes.InvalidArgument, "order not found: %v", err)
    }

    if orderResp.Order.Status != "confirmed" {
        return nil, status.Errorf(codes.FailedPrecondition, "order is not confirmed yet")
    }

    s.mu.Lock()
    defer s.mu.Unlock()

    reviewId := fmt.Sprintf("review_%d", len(s.reviews)+1)

    review := &reviewPb.Review{
        Id:        reviewId,
        UserId:    req.UserId,
        OrderId:   req.OrderId,
        Rating:    req.Rating,
        Comment:   req.Comment,
        CreatedAt: time.Now().Format(time.RFC3339),
    }

    s.reviews[reviewId] = review

    return &reviewPb.ReviewResponse{Review: review}, nil
}

func (s *reviewService) GetReview(ctx context.Context, req *reviewPb.GetReviewRequest) (*reviewPb.ReviewResponse, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    review, exists := s.reviews[req.Id]
    if !exists {
        return nil, status.Errorf(codes.NotFound, "review not found")
    }

    return &reviewPb.ReviewResponse{Review: review}, nil
}

func main() {
    userConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to connect to user service: %v", err)
    }
    defer userConn.Close()

    orderConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to connect to order service: %v", err)
    }
    defer orderConn.Close()

    userClient := userPb.NewUserServiceClient(userConn)
    orderClient := orderPb.NewOrderServiceClient(orderConn)

    lis, err := net.Listen("tcp", ":50054")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    reviewPb.RegisterReviewServiceServer(grpcServer, newReviewService(userClient, orderClient))

    log.Println("Starting review service on :50054")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
