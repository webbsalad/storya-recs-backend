package pg

import "github.com/webbsalad/storya-recs-backend/internal/model"

func ValueToDeviation(value model.Value) int {
	switch value {
	case 0:
		return 1
	case 1:
		return 0
	case 2:
		return -1
	default:
		return 0
	}
}
