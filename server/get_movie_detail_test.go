package server

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/muhammadarash1997/movies/model"
)

type GetMovieDetailTestCase struct {
	TestingName       string
	Request           int64
	WantResponse      model.Movie
	WantError         bool
	ReadMovieResponse model.Movie
	ReadMovieError    error
}

func getGetMovieDetailTestCases(t *testing.T) []GetMovieDetailTestCase {
	return []GetMovieDetailTestCase{
		{
			TestingName: "Passed",
			Request:     1,
			WantResponse: model.Movie{
				ID:          1,
				Title:       "A",
				Description: "B",
				Rating:      1,
				Image:       "ABC",
				CreatedAt:   timeNowDummy,
				UpdatedAt:   timeNowDummy,
			},
			WantError: false,
			ReadMovieResponse: model.Movie{
				ID:          1,
				Title:       "A",
				Description: "B",
				Rating:      1,
				Image:       "ABC",
				CreatedAt:   timeNowDummy,
				UpdatedAt:   timeNowDummy,
			},
			ReadMovieError: nil,
		},
	}
}

func Test_GetMovieDetail(t *testing.T) {
	initTest(t)
	for _, v := range getGetMovieDetailTestCases(t) {
		t.Run(v.TestingName, func(t *testing.T) {
			db.EXPECT().ReadMovie(gomock.Any()).Return(v.ReadMovieResponse, v.ReadMovieError).AnyTimes()

			gotResp, err := serviceDummy.GetMovieDetail(context.Background(), v.Request)
			if (err != nil) != v.WantError {
				t.Errorf("MoviesService.GetMovieDetail() Error = %v, WantError %v", err, v.WantError)
				return
			}
			if !reflect.DeepEqual(gotResp, v.WantResponse) {
				t.Errorf("MoviesService.GetMovieDetail() Response = %v, WantResponse %v", gotResp, v.WantResponse)
			}
		})
	}
}
