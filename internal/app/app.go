package app

import (
	"github.com/webbsalad/storya-recs-backend/internal/api/recs"
	"github.com/webbsalad/storya-recs-backend/internal/config"
	pb "github.com/webbsalad/storya-recs-backend/internal/pb/github.com/webbsalad/storya-recs-backend/recs"
	"github.com/webbsalad/storya-recs-backend/internal/repository/recs/pg"

	v1 "github.com/webbsalad/storya-recs-backend/internal/service/recs/v1"
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			config.NewConfig,
			initDB,
		),
		grpcOptions(),
		serviceOptions(),
	)
}

func serviceOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			recs.NewRecsImplementation,
			pg.NewRepository,
			v1.NewService,
		),
		fx.Invoke(
			pb.RegisterRecsServiceServer,
		),
	)
}
