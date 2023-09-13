package server

import (
	"context"
	"reflect"
	"testing"

	"github.com/muhammadarash1997/movies/model"
)

type GetMovieListTestCase struct {
	TestingName           string
	WantResponse          []model.Movie
	WantError             bool
	ReadAllMoviesResponse []model.Movie
	ReadAllMoviesError    error
}

func getGetMovieListTestCases(t *testing.T) []GetMovieListTestCase {
	return []GetMovieListTestCase{
		{
			TestingName: "Passed",
			WantResponse: []model.Movie{
				{
					ID:          1,
					Title:       "A",
					Description: "B",
					Rating:      1,
					Image:       "ABC",
				},
				{
					ID:          2,
					Title:       "B",
					Description: "C",
					Rating:      2,
					Image:       "ABCD",
				},
			},
			WantError: false,
			ReadAllMoviesResponse: []model.Movie{
				{
					ID:          1,
					Title:       "A",
					Description: "B",
					Rating:      1,
					Image:       "ABC",
				},
				{
					ID:          2,
					Title:       "B",
					Description: "C",
					Rating:      2,
					Image:       "ABCD",
				},
			},
			ReadAllMoviesError: nil,
		},
	}
}

func Test_GetMovieList(t *testing.T) {
	initTest(t)
	for _, v := range getGetMovieListTestCases(t) {
		t.Run(v.TestingName, func(t *testing.T) {
			db.EXPECT().ReadAllMovies().Return(v.ReadAllMoviesResponse, v.ReadAllMoviesError).AnyTimes()

			gotResp, err := serviceDummy.GetMovieList(context.Background())
			if (err != nil) != v.WantError {
				t.Errorf("MoviesService.GetMovieList() Error = %v, WantError %v", err, v.WantError)
				return
			}
			if !reflect.DeepEqual(gotResp, v.WantResponse) {
				t.Errorf("MoviesService.GetMovieList() Response = %v, WantResponse %v", gotResp, v.WantResponse)
			}
		})
	}
}
