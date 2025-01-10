package recs

import (
	"context"

	"github.com/webbsalad/storya-recs-backend/internal/model"
)

type Repository interface {
	GetPreferences(ctx context.Context, userID model.UserID) ([]model.Preference, error)
	UpdatePreferences(ctx context.Context, userID model.UserID, ratedTags []model.RatedTag) ([]model.Preference, error)
	DeletePeferences(ctx context.Context, userID model.UserID) error
}
