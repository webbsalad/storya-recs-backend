package recs

import (
	desc "github.com/webbsalad/storya-recs-backend/internal/pb/github.com/webbsalad/storya-recs-backend/recs"
	service "github.com/webbsalad/storya-recs-backend/internal/service/recs"
)

type Implementation struct {
	desc.UnimplementedRecsServiceServer

	RecsService service.Service
}

func NewRecsImplementation(recsService service.Service) desc.RecsServiceServer {
	return &Implementation{
		RecsService: recsService,
	}
}
