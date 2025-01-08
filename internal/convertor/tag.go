package convertor

import (
	"github.com/webbsalad/storya-recs-backend/internal/model"
	"github.com/webbsalad/storya-recs-backend/internal/pb/github.com/webbsalad/storya-recs-backend/recs"
)

func ToDescFromTags(in []model.Tag) []*recs.Tag {
	tags := make([]*recs.Tag, len(in))

	for i, tg := range in {
		tags[i] = ToDescFromTag(tg)
	}

	return tags
}

func ToDescFromTag(tg model.Tag) *recs.Tag {
	return &recs.Tag{
		Name: tg.Name,
	}
}

func ToTagsFromDesc(in []*recs.Tag) []model.Tag {
	tags := make([]model.Tag, len(in))

	for i, tg := range in {
		tags[i] = toTagFromDesc(tg)
	}

	return tags
}

func toTagFromDesc(tg *recs.Tag) model.Tag {
	return model.Tag{
		Name: tg.Name,
	}
}
