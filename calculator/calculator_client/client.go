package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Mark-Jung/grpc-go/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("me client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Doing Unary...")

	req := &calculatorpb.SumRequest{
		Sum: &calculatorpb.Sum{
			First:  4,
			Second: 9,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum rpc: %v", err)
	}
	log.Printf("Response from Calculator rpc: %v", res.Result)
}
