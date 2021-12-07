package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/ditdittdittt/learn-gql/graph/generated"
	"github.com/ditdittdittt/learn-gql/graph/model"
	pb "github.com/ditdittdittt/learn-grpc/productproto"
)

func (r *mutationResolver) Create(ctx context.Context, name string) (*model.CreateResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	res, err := r.ProductClient.Create(ctx, &pb.CreateRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}

	resp := &model.CreateResponse{
		ID:           res.GetID(),
		ResponseCode: int(res.GetResponseCode()),
	}

	return resp, nil
}

func (r *mutationResolver) Update(ctx context.Context, id string, name string) (*model.UpdateResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	res, err := r.ProductClient.Update(ctx, &pb.UpdateRequest{
		Name: name,
		ID:   id,
	})
	if err != nil {
		return nil, err
	}

	resp := &model.UpdateResponse{
		ID:           res.GetID(),
		Name:         res.GetName(),
		ResponseCode: int(res.GetResponseCode()),
	}

	return resp, nil
}

func (r *mutationResolver) Delete(ctx context.Context, id string) (*model.DeleteResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	res, err := r.ProductClient.Delete(ctx, &pb.DeleteRequest{
		ID: id,
	})
	if err != nil {
		return nil, err
	}

	resp := &model.DeleteResponse{
		ResponseCode: int(res.GetResponseCode()),
	}

	return resp, nil
}

func (r *queryResolver) Read(ctx context.Context) (*model.ReadResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	res, err := r.ProductClient.Read(ctx, &pb.ReadRequest{})
	if err != nil {
		return nil, err
	}

	resp := &model.ReadResponse{
		ResponseCode: int(res.GetResponseCode()),
	}
	for _, product := range res.GetProducts() {
		resp.Products = append(resp.Products, &model.Product{
			ID:   product.ID,
			Name: product.Name,
		})
	}

	return resp, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
