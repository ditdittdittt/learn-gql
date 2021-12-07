package graph

import (
	pb "github.com/ditdittdittt/learn-grpc/productproto"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProductClient pb.ProductServiceClient
}
