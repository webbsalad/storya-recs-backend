package convertor

import (
	"github.com/webbsalad/storya-recs-backend/internal/model"
	"github.com/webbsalad/storya-recs-backend/internal/pb/github.com/webbsalad/storya-recs-backend/recs"
)

func ToDescsFromPreferences(in []model.Preference) []*recs.Preference {
	preferences := make([]*recs.Preference, len(in))
	for i, preference := range in {
		preferences[i] = ToDescFrompreference(preference)
	}

	return preferences

}

func ToDescFrompreference(in model.Preference) *recs.Preference {
	return &recs.Preference{
		Tag:   ToDescFromTag(in.Tag),
		Value: in.Value,
	}
}
