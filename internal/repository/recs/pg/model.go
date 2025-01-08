package pg

import (
	"time"

	"github.com/webbsalad/storya-recs-backend/internal/model"
)

type Tag struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

type Preferences struct {
	UserID    string    `db:"user_id"`
	TagID     string    `db:"tag_id"`
	Value     int32     `db:"value"`
	UpdatedAt time.Time `db:"updated_at"`
}

type FullRatedTag struct {
	ID string `db:"id"`

	RatedTag model.RatedTag
}
