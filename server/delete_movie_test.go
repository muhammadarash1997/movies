package server

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
)

type DeleteMovieTestCase struct {
	TestingName      string
	Request          int64
	WantError        bool
	DeleteMovieError error
}

func getDeleteMovieTestCases(t *testing.T) []DeleteMovieTestCase {
	return []DeleteMovieTestCase{
		{
			TestingName:      "Passed",
			Request:          11,
			WantError:        false,
			DeleteMovieError: nil,
		},
	}
}

func Test_DeleteMovie(t *testing.T) {
	initTest(t)
	for _, v := range getDeleteMovieTestCases(t) {
		t.Run(v.TestingName, func(t *testing.T) {
			db.EXPECT().DeleteMovie(gomock.Any()).Return(v.DeleteMovieError).AnyTimes()

			err := serviceDummy.DeleteMovie(context.Background(), v.Request)
			if (err != nil) != v.WantError {
				t.Errorf("MoviesService.DeleteMovie() Error = %v, WantError %v", err, v.WantError)
				return
			}
		})
	}
}
