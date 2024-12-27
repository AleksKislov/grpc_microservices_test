package main

import (
    "context"
    "log"
    "net"
    "sync"

    pb "user-service/proto/user"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "github.com/google/uuid"
)

type server struct {
    pb.UnimplementedUserServiceServer
    users map[string]*pb.User
    mutex sync.RWMutex
}

func newServer() *server {
    return &server{
        users: make(map[string]*pb.User),
    }
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    // Check if email already exists
    for _, user := range s.users {
        if user.Email == req.Email {
            return nil, status.Errorf(codes.AlreadyExists, "user with email %s already exists", req.Email)
        }
    }

    user := &pb.User{
        Id:    uuid.New().String(),
        Email: req.Email,
        Name:  req.Name,
        Phone: req.Phone,
    }

    s.users[user.Id] = user

    return &pb.UserResponse{
        User: user,
    }, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
    s.mutex.RLock()
    defer s.mutex.RUnlock()

    user, exists := s.users[req.Id]
    if !exists {
        return nil, status.Errorf(codes.NotFound, "user not found")
    }

    return &pb.UserResponse{
        User: user,
    }, nil
}

func (s *server) AuthenticateUser(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
    s.mutex.RLock()
    defer s.mutex.RUnlock()

    // Simple authentication (in real world, you'd hash passwords and use proper auth)
    for _, user := range s.users {
        if user.Email == req.Email {
            // Generate a simple token (in real world, use proper JWT)
            token := "dummy-token-" + user.Id
            return &pb.AuthResponse{
                Token: token,
                User:  user,
            }, nil
        }
    }

    return nil, status.Errorf(codes.NotFound, "user not found")
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    
    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, newServer())
    
    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
