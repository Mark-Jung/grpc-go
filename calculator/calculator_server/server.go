package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Mark-Jung/grpc-go/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v", req)
	firstNum := req.GetSum().GetFirst()
	secondNum := req.GetSum().GetSecond()
	sum := firstNum + secondNum
	res := &calculatorpb.SumResponse{
		Result: sum,
	}
	return res, nil
}

func main() {
	fmt.Println("Hi")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
