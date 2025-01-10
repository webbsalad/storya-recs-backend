package v1

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/webbsalad/storya-recs-backend/internal/config"
	"github.com/webbsalad/storya-recs-backend/internal/model"
	mock_repository "github.com/webbsalad/storya-recs-backend/internal/repository/recs/mock"
)

var (
	testUserID, _ = model.UserIDFromString("2b98ee88-7970-4e6f-b325-ccf3ce10909f")
)

type serviceTestDeps struct {
	Service *Service

	ctx            context.Context
	recsRepository *mock_repository.MockRepository
}

func getTestDeps(t *testing.T) *serviceTestDeps {
	ctrl := gomock.NewController(t)
	recsRepository := mock_repository.NewMockRepository(ctrl)

	return &serviceTestDeps{
		Service: &Service{
			recsRepository: recsRepository,
			config:         config.Config{},
		},

		ctx:            context.Background(),
		recsRepository: recsRepository,
	}
}
