package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"goapi/graph/generated"
	"goapi/graph/model"
	"goapi/service"
)

func (r *authOpsResolver) Login(ctx context.Context, obj *model.AuthOps, email string, password string) (interface{}, error) {
	return service.UserLogin(ctx, email, password)
}

func (r *authOpsResolver) Register(ctx context.Context, obj *model.AuthOps, input model.InputUser) (interface{}, error) {
	return service.UserRegister(ctx, input)
}

func (r *mutationResolver) Auth(ctx context.Context) (*model.AuthOps, error) {
	return &model.AuthOps{}, nil
}

func (r *mutationResolver) AddCategory(ctx context.Context, category model.InputCategory) (interface{}, error) {
	return service.AddCategory(ctx, category)
}

func (r *mutationResolver) EditCategory(ctx context.Context, category model.InputCategory, categoryID int64) (interface{}, error) {
	return service.EditCategory(ctx, category, categoryID)
}

func (r *mutationResolver) DeleteCategory(ctx context.Context, id int64) (interface{}, error) {
	return service.DeleteCategory(ctx, id)
}

func (r *mutationResolver) AddProduct(ctx context.Context, product model.InputProduct) (interface{}, error) {
	return service.AddProduct(ctx, product)
}

func (r *mutationResolver) EditProduct(ctx context.Context, product model.InputProduct, productID int64) (interface{}, error) {
	return service.EditProduct(ctx, product, productID)
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, id int64) (interface{}, error) {
	return service.DeleteProduct(id)
}

func (r *queryResolver) User(ctx context.Context, id int64) (*model.User, error) {
	return service.UserGetByID(ctx, id)
}

func (r *queryResolver) Protected(ctx context.Context) (string, error) {
	return "Success", nil
}

func (r *queryResolver) Products(ctx context.Context, search *string, categoryID *int64) ([]*model.Product, error) {
	return service.Products(ctx, search, categoryID)
}

func (r *queryResolver) Categories(ctx context.Context, search *string) ([]*model.Category, error) {
	return service.Categories(ctx, search)
}

// AuthOps returns generated.AuthOpsResolver implementation.
func (r *Resolver) AuthOps() generated.AuthOpsResolver { return &authOpsResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type authOpsResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
