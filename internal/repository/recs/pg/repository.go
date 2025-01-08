package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/webbsalad/storya-recs-backend/internal/model"
	"github.com/webbsalad/storya-recs-backend/internal/repository/recs"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) (recs.Repository, error) {
	return &Repository{db: db}, nil
}

func (r *Repository) UpdatePreferences(ctx context.Context, userID model.UserID, ratedTags []model.RatedTag) ([]model.Preference, error) {
	var tagsData []FullRatedTag

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		rollbackErr := tx.Rollback()
		if err == nil && rollbackErr != nil {
			err = fmt.Errorf("rollback: %w", rollbackErr)
		}
	}()

	for _, ratedTag := range ratedTags {
		var tagData FullRatedTag

		tagsSelectQuery := psql.
			Select("id").
			From("tag").
			Where(
				sq.Eq{"name": ratedTag.Tag.Name},
			)

		q, args, err := tagsSelectQuery.ToSql()
		if err != nil {
			return nil, fmt.Errorf("build select tag query: %w", err)
		}

		if err = tx.GetContext(ctx, &tagData, q, args...); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, model.ErrTagNotFound
			}
			return nil, fmt.Errorf("select tag: %w", err)
		}

		tagData.RatedTag = ratedTag
		tagsData = append(tagsData, tagData)

	}

	for _, tagData := range tagsData {
		insertPreferencesQuery := psql.
			Insert("user_preferences").
			Columns("user_id", "tag_id", "value", "updated_at").
			Values(userID.String(), tagData.ID, ValueToDeviation(tagData.RatedTag.Value), time.Now()).
			Suffix("ON CONFLICT (user_id, tag_id) DO UPDATE SET value = user_preferences.value + EXCLUDED.value, updated_at = EXCLUDED.updated_at")

		q, args, err := insertPreferencesQuery.ToSql()
		if err != nil {
			return nil, fmt.Errorf("insert preferences query: %w", err)
		}

		_, err = tx.ExecContext(ctx, q, args...)
		if err != nil {
			return nil, fmt.Errorf("execute insert preferences: %w", err)
		}
	}

	var newPreferences []Preferences

	selectTagIDsQuery := psql.
		Select("tag_id", "value").
		From("user_preferences").
		Where(
			sq.Eq{"user_id": userID.String()},
		)

	q, args, err := selectTagIDsQuery.ToSql()
	if err != nil {
		return nil, fmt.Errorf("select tag ids query: %w", err)
	}

	if err = tx.SelectContext(ctx, &newPreferences, q, args...); err != nil {
		return nil, fmt.Errorf("exec select tag ids query: %w", err)
	}

	updatedPreferences := make([]model.Preference, len(newPreferences))
	for i, newPreference := range newPreferences {
		var tag Tag
		selectTagNamesQuery := psql.
			Select("name").
			From("tag").
			Where(
				sq.Eq{"id": newPreference.TagID},
			)

		q, args, err := selectTagNamesQuery.ToSql()
		if err != nil {
			return nil, fmt.Errorf("build select tag names query: %w", err)
		}

		if err = tx.GetContext(ctx, &tag, q, args...); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, model.ErrTagNotFound
			}
			return nil, fmt.Errorf("select tag name: %w", err)
		}

		updatedPreferences[i].Tag.Name = tag.Name
		updatedPreferences[i].Value = newPreference.Value
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("commit transaction: %w", err)
	}

	return updatedPreferences, nil
}
