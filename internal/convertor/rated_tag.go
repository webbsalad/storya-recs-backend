package convertor

import (
	"github.com/webbsalad/storya-recs-backend/internal/model"
	"github.com/webbsalad/storya-recs-backend/internal/pb/github.com/webbsalad/storya-recs-backend/recs"
)

func ToRatedTagsFromDescs(in []*recs.RatedTag) []model.RatedTag {
	ratedTags := make([]model.RatedTag, len(in))
	for i, ratedTag := range in {
		ratedTags[i] = ToRatedTagFromDesc(ratedTag)
	}

	return ratedTags
}

func ToRatedTagFromDesc(in *recs.RatedTag) model.RatedTag {
	return model.RatedTag{
		Tag:   toTagFromDesc(in.Tag),
		Value: model.Value(in.Value),
	}
}

func ToDescsFromRatedTags(in []model.RatedTag) []*recs.RatedTag {
	ratedTags := make([]*recs.RatedTag, len(in))
	for i, rateratedTag := range in {
		ratedTags[i] = ToDescFromRatedTag(rateratedTag)
	}

	return ratedTags

}

func ToDescFromRatedTag(in model.RatedTag) *recs.RatedTag {
	return &recs.RatedTag{
		Tag:   ToDescFromTag(in.Tag),
		Value: recs.Value(in.Value),
	}
}
