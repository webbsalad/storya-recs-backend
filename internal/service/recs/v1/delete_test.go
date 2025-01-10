package v1

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/webbsalad/storya-recs-backend/internal/model"
	"go.uber.org/mock/gomock"
)

func TestService_DeletePreferences(t *testing.T) {
	t.Parallel()

	type args struct {
		userID model.UserID
	}

	type result struct {
		err error
	}

	type testCase struct {
		name   string
		args   args
		mocks  func(tc testCase, deps *serviceTestDeps)
		result result
	}

	testCases := []testCase{
		{
			name: "success",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				deps.recsRepository.EXPECT().
					DeletePeferences(gomock.Any(), tc.args.userID).
					Return(nil)
			},
			args: args{
				userID: testUserID,
			},
			result: result{
				err: nil,
			},
		},

		{
			name: "user not found",
			mocks: func(tc testCase, deps *serviceTestDeps) {
				deps.recsRepository.EXPECT().
					DeletePeferences(gomock.Any(), tc.args.userID).
					Return(model.ErrPreferencesNotFound)
			},
			args: args{
				userID: testUserID,
			},
			result: result{
				err: model.ErrPreferencesNotFound,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			deps := getTestDeps(t)

			tc.mocks(tc, deps)
			err := deps.Service.DeletePeferences(deps.ctx, tc.args.userID)

			require.ErrorIs(t, err, tc.result.err)

		})
	}
}
