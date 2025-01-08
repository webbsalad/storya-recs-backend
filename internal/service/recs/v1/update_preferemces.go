package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/storya-recs-backend/internal/model"
)

func (s *Service) UpdatePreferences(ctx context.Context, userID model.UserID, ratedTags []model.RatedTag) ([]model.Preference, error) {
	updatedPreferences, err := s.recsRepository.UpdatePreferences(ctx, userID, ratedTags)
	if err != nil {
		return nil, fmt.Errorf("update: %w", err)
	}

	return updatedPreferences, nil
}
