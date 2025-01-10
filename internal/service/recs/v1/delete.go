package v1

import (
	"context"
	"fmt"

	"github.com/webbsalad/storya-recs-backend/internal/model"
)

func (s *Service) DeletePeferences(ctx context.Context, userID model.UserID) error {
	if err := s.recsRepository.DeletePeferences(ctx, userID); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}
