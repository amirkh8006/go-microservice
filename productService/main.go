package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/amirkh8006/go-microservice/productpb"
	"google.golang.org/grpc"
)

type Product struct {
	ID int
	Name string
	UsdPerUnit float64
	Unit string
}

var products = []Product{
	{
		ID:         1,
		Name:       "Product 1",
		UsdPerUnit: 10.0,
		Unit:       "kg",
	},
	{
		ID:         2,
		Name:       "Product 2",
		UsdPerUnit: 20.0,
		Unit:       "kg",
	},
	{
		ID:         3,
		Name:       "Product 3",
		UsdPerUnit: 30.0,
		Unit:       "Pound",
	},
	{
		ID:         4,
		Name:       "Product 4",
		UsdPerUnit: 40.0,
		Unit:       "Pound",
	},
}

func main() {
	go startGrpcServer()
	log.Println("gRPC server started")

	// DONOT USE time.Sleep in production code, this is just for demonstration purposes
	time.Sleep(1 * time.Second)

	callGRPCService()
}

type ProductService struct {
	productpb.UnimplementedProductServer
}

func (ps ProductService) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.GetProductResponse, error) {
	log.Println("Received GetProduct request for ID:", req.Id)

	for _, product := range products {
		if int32(product.ID) == req.Id {
			return &productpb.GetProductResponse{
				Product: &productpb.Product{
					Id:         int32(product.ID),
					Name:       product.Name,
					UsdPerUnit: product.UsdPerUnit,
					Unit:       product.Unit,
				},
			}, nil
		}
	}

	return nil, fmt.Errorf("product with ID %d not found", req.Id)
}


func startGrpcServer() {
	lis , err := net.Listen("tcp", ":4001")
	if err != nil {
		log.Fatalf("Failed to listen on port 4001: %v", err)
	}

	grpcServer := grpc.NewServer()
	productpb.RegisterProductServer(grpcServer, &ProductService{})
	log.Println("Starting gRPC server on port 4001...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

func callGRPCService() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial("localhost:4001", opts...)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()
	client := productpb.NewProductClient(conn)
	res, err := client.GetProduct(context.Background(), &productpb.GetProductRequest{
		Id: 1,
	})

	if err != nil {
		log.Fatalf("Error calling GetProduct: %v", err)
	}

	log.Printf("Received response: %v", res.Product)
}