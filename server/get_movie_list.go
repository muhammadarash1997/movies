package server

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/log/level"
	"github.com/muhammadarash1997/movies/model"
	shv "github.com/muhammadarash1997/movies/utils/sharevar"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Movies) GetMovieList(ctx context.Context) (movies []model.Movie, err error) {
	const (
		funcName = `GetMovieList`
	)

	timeStart := time.Now()
	level.Info(shv.Logger).Log("[INFO]", fmt.Sprintf("Higher %s", funcName))

	movies, err = s.repo.DBReadWriter.ReadAllMovies()
	if err != nil {
		level.Error(shv.Logger).Log("[ERROR]", err.Error())
		return movies, status.New(codes.Internal, codes.Internal.String()).Err()
	}

	level.Info(shv.Logger).Log("[INFO]", fmt.Sprintf("Lower %s", funcName), "[DURATION]", time.Since(timeStart).Seconds(), "[RESPONSE]", fmt.Sprintf("%+v", movies))
	return movies, nil
}
