package recs

import (
	"context"

	"github.com/webbsalad/storya-recs-backend/internal/convertor"
	"github.com/webbsalad/storya-recs-backend/internal/model"
	"github.com/webbsalad/storya-recs-backend/internal/pb/github.com/webbsalad/storya-recs-backend/recs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetPreferences(ctx context.Context, req *recs.GetPreferencesRequest) (*recs.GetPreferencesResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	userID, err := model.UserIDFromString(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid data: %v", err)
	}

	preferences, err := i.RecsService.GetPreferences(ctx, userID)
	if err != nil {
		return nil, convertor.ConvertError(err)
	}

	return &recs.GetPreferencesResponse{
		Preferences: convertor.ToDescsFromPreferences(preferences),
	}, nil

}
