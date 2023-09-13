package server

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/muhammadarash1997/movies/model"
)

type UpdateMovieTestCase struct {
	TestingName             string
	Request                 model.Movie
	WantError               bool
	CheckMovieExistResponse bool
	CheckMovieExistError    error
	UpdateMovieError        error
}

func getUpdateMovieTestCases(t *testing.T) []UpdateMovieTestCase {
	return []UpdateMovieTestCase{
		{
			TestingName: "Passed",
			Request: model.Movie{
				ID:          1,
				Title:       "A",
				Description: "B",
				Rating:      1,
				Image:       "ABC",
				CreatedAt:   timeNowDummy,
				UpdatedAt:   timeNowDummy,
			},
			WantError:               false,
			CheckMovieExistResponse: true,
			CheckMovieExistError:    nil,
			UpdateMovieError:        nil,
		},
	}
}

func Test_UpdateMovie(t *testing.T) {
	initTest(t)
	for _, v := range getUpdateMovieTestCases(t) {
		t.Run(v.TestingName, func(t *testing.T) {
			db.EXPECT().CheckMovieExist(gomock.Any()).Return(v.CheckMovieExistResponse, v.CheckMovieExistError).AnyTimes()
			db.EXPECT().UpdateMovie(gomock.Any()).Return(v.UpdateMovieError).AnyTimes()

			err := serviceDummy.UpdateMovie(context.Background(), v.Request)
			if (err != nil) != v.WantError {
				t.Errorf("MoviesService.UpdateMovie() Error = %v, WantError %v", err, v.WantError)
				return
			}
		})
	}
}
