package server

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/muhammadarash1997/movies/model"
)

type AddMovieTestCase struct {
	TestingName     string
	Request         model.Movie
	WantError       bool
	WriteMovieError error
}

func getAddMovieTestCases(t *testing.T) []AddMovieTestCase {
	return []AddMovieTestCase{
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
			WantError:       false,
			WriteMovieError: nil,
		},
	}
}

func Test_AddMovie(t *testing.T) {
	initTest(t)
	for _, v := range getAddMovieTestCases(t) {
		t.Run(v.TestingName, func(t *testing.T) {
			db.EXPECT().WriteMovie(gomock.Any()).Return(v.WriteMovieError).AnyTimes()

			err := serviceDummy.AddMovie(context.Background(), v.Request)
			if (err != nil) != v.WantError {
				t.Errorf("MoviesService.AddMovie() Error = %v, WantError %v", err, v.WantError)
				return
			}
		})
	}
}
