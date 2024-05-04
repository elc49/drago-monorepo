package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"strconv"
	"strings"

	"github.com/edwinlomolo/drago-api/graph/model"
	"github.com/google/uuid"
)

// CreateBusiness is the resolver for the createBusiness field.
func (r *mutationResolver) CreateBusiness(ctx context.Context, input model.NewBusinessInput) (*model.Business, error) {
	userId := ctx.Value("userId").(string)
	ip := ctx.Value("ip").(string)
	location, err := r.ip.GetIpinfo(ip)
	if err != nil {
		return nil, err
	}

	uId, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}
	input.UserID = uId

	coords := strings.Split(location.Location, ",")
	lat, err := strconv.ParseFloat(coords[0], 64)
	if err != nil {
		return nil, err
	}

	lng, err := strconv.ParseFloat(coords[1], 64)
	if err != nil {
		return nil, err
	}

	input.Location = model.Gps{Lat: lat, Lng: lng}
	return r.bc.CreateBusiness(ctx, input)
}

// AddCourier is the resolver for the addCourier field.
func (r *mutationResolver) AddCourier(ctx context.Context, input model.NewCourierInput) (*model.Courier, error) {
	return r.bc.CreateBusinessCourier(ctx, input)
}

// CreateTrip is the resolver for the createTrip field.
func (r *mutationResolver) CreateTrip(ctx context.Context, input model.NewTrip) (*model.Trip, error) {
	pickup, err := r.l.GetPlaceDetails(ctx, input.Pickup)
	if err != nil {
		return nil, err
	}

	dropoff, err := r.l.GetPlaceDetails(ctx, input.Dropoff)
	if err != nil {
		return nil, err
	}

	args := model.NewTripInput{
		PickupAddress:  pickup.FormattedAddress,
		DropoffAddress: dropoff.FormattedAddress,
		Pickup:         pickup.Coords,
		Dropoff:        dropoff.Coords,
		CourierID:      input.CourierID,
		BusinessID:     input.BusinessID,
	}
	return r.tc.CreateTrip(ctx, args)
}

// SetUserDefaultBusiness is the resolver for the setUserDefaultBusiness field.
func (r *mutationResolver) SetUserDefaultBusiness(ctx context.Context, businessID uuid.UUID) (*model.User, error) {
	userId := ctx.Value("userId").(string)
	uId, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return r.uc.SetDefaultBusiness(ctx, uId, businessID)
}

// GetBusinessBelongingToUser is the resolver for the getBusinessBelongingToUser field.
func (r *queryResolver) GetBusinessBelongingToUser(ctx context.Context) ([]*model.Business, error) {
	userId := ctx.Value("userId").(string)
	uId, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return r.bc.GetBusinessBelongingToUser(ctx, uId)
}

// GetBusinessCouriers is the resolver for the getBusinessCouriers field.
func (r *queryResolver) GetCouriersBelongingToBusiness(ctx context.Context) ([]*model.Courier, error) {
	userId := ctx.Value("userId").(string)
	uId, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return r.bc.GetBusinessCouriers(ctx, uId)
}

// GetBusinessDeliveryTrips is the resolver for the getBusinessDeliveryTrips field.
func (r *queryResolver) GetTripsBelongingToBusiness(ctx context.Context) ([]*model.Trip, error) {
	userId := ctx.Value("userId").(string)
	uId, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return r.tc.GetTripsBelongingToBusiness(ctx, uId)
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context) (*model.User, error) {
	userId := ctx.Value("userId").(string)
	uId, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return r.uc.GetUserByID(ctx, uId)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
