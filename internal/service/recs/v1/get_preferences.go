package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/storya-recs-backend/internal/model"
)

func (s *Service) GetPreferences(ctx context.Context, userID model.UserID) ([]model.Preference, error) {
	preferences, err := s.recsRepository.GetPreferences(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return preferences, nil
}
