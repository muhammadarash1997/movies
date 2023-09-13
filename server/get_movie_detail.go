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

func (s *Movies) GetMovieDetail(ctx context.Context, id int64) (movie model.Movie, err error) {
	const (
		funcName = `GetMovieDetail`
	)

	timeStart := time.Now()
	level.Info(shv.Logger).Log("[INFO]", fmt.Sprintf("Higher %s", funcName), "[REQUEST]", fmt.Sprintf("%+v", id))

	if id == 0 {
		level.Error(shv.Logger).Log("[ERROR]", "Invalid request")
		return movie, status.New(codes.InvalidArgument, codes.InvalidArgument.String()).Err()
	}

	movie, err = s.repo.DBReadWriter.ReadMovie(id)
	if err != nil {
		level.Error(shv.Logger).Log("[ERROR]", err.Error())
		return movie, status.New(codes.Internal, codes.Internal.String()).Err()
	}

	if movie.IsEmpty() {
		level.Info(shv.Logger).Log("[INFO]", "Data is not exist")
		return movie, status.New(codes.NotFound, codes.NotFound.String()).Err()
	}

	level.Info(shv.Logger).Log("[INFO]", fmt.Sprintf("Lower %s", funcName), "[DURATION]", time.Since(timeStart).Seconds(), "[RESPONSE]", fmt.Sprintf("%+v", movie))
	return movie, nil
}

func idRequestValidator(id int64) bool {
	return id != 0
}
