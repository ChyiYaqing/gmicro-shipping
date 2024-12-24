package grpc

import (
	"context"
	"fmt"

	"github.com/chyiyaqing/gmicro-proto/golang/shipping"
	"github.com/chyiyaqing/gmicro-shipping/internal/application/core/domain"
	log "github.com/sirupsen/logrus"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Adapter) Create(ctx context.Context, request *shipping.CreateShippingRequest) (*shipping.CreateShippingResponse, error) {
	log.WithContext(ctx).Info("Creating shipping...")
	var validationErrors []*errdetails.BadRequest_FieldViolation
	if request.UserId < 1 {
		validationErrors = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field:       "user_id",
			Description: "user id canot be less than 1",
		})
	}
	if len(validationErrors) > 0 {
		stat := status.New(400, "invalid order request")
		badRequest := &errdetails.BadRequest{}
		badRequest.FieldViolations = validationErrors
		s, _ := stat.WithDetails(badRequest)
		return nil, s.Err()
	}
	newShipping := domain.NewShipping(request.UserId, request.OrderId, request.Address)
	result, err := a.api.Create(ctx, newShipping)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v", err)).Err()
	}
	return &shipping.CreateShippingResponse{ShippingId: result.ID}, nil
}
