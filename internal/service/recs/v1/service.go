package v1

import (
	"github.com/webbsalad/storya-recs-backend/internal/config"
	"github.com/webbsalad/storya-recs-backend/internal/repository/recs"
	recs_service "github.com/webbsalad/storya-recs-backend/internal/service/recs"
)

type Service struct {
	recsRepository recs.Repository
	config         config.Config
}

func NewService(recsRepository recs.Repository, config config.Config) recs_service.Service {
	return &Service{
		recsRepository: recsRepository,
		config:         config,
	}
}
