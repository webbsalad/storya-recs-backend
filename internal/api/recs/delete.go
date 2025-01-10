package recs

import (
	"context"

	"github.com/webbsalad/storya-recs-backend/internal/convertor"
	"github.com/webbsalad/storya-recs-backend/internal/utils/metadata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) DeletePeferences(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	userID, err := metadata.GetUserID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid data: %v", err)
	}

	if err := i.RecsService.DeletePeferences(ctx, userID); err != nil {
		return nil, convertor.ConvertError(err)
	}

	return &emptypb.Empty{}, nil
}
