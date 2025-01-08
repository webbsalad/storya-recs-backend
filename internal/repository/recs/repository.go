package recs

import (
	"context"

	"github.com/webbsalad/storya-recs-backend/internal/model"
)

type Repository interface {
	UpdatePreferences(ctx context.Context, userID model.UserID, ratedTags []model.RatedTag) ([]model.Preference, error)
}
