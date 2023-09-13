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

func (s *Movies) AddMovie(ctx context.Context, request model.Movie) error {
	const (
		funcName = `AddMovie`
	)

	timeStart := time.Now()
	level.Info(shv.Logger).Log("[INFO]", fmt.Sprintf("Higher %s", funcName), "[REQUEST]", fmt.Sprintf("%+v", request))

	if !request.Validate() {
		level.Error(shv.Logger).Log("[ERROR]", "Invalid request")
		return status.New(codes.InvalidArgument, codes.InvalidArgument.String()).Err()
	}

	err := s.repo.DBReadWriter.WriteMovie(request)
	if err != nil {
		level.Error(shv.Logger).Log("[ERROR]", err.Error())
		return status.New(codes.Internal, codes.Internal.String()).Err()
	}

	level.Info(shv.Logger).Log("[INFO]", fmt.Sprintf("Lower %s", funcName), "[DURATION]", time.Since(timeStart).Seconds())
	return nil
}