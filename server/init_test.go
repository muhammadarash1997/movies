package server

import (
	"os"
	"testing"

	"time"

	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	"github.com/muhammadarash1997/movies/repository"
	repoMock "github.com/muhammadarash1997/movies/repository/interface"
	"github.com/muhammadarash1997/movies/service"
	shv "github.com/muhammadarash1997/movies/utils/sharevar"
)

var (
	db           *repoMock.MockReadWriter
	serverDummy  *Movies
	serviceDummy service.MoviesService

	timeNowDummy time.Time
)

func initTest(t *testing.T) {
	timeNowDummy = time.Now()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	shv.Logger = log.NewLogfmtLogger(os.Stdout)
	shv.Logger = log.With(shv.Logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	db = repoMock.NewMockReadWriter(mockCtrl)

	serverDummy = &Movies{
		repo: repository.Repo{
			DBReadWriter: db,
		},
	}
	serviceDummy = NewMoviesService(repository.Repo{
		DBReadWriter: db,
	})
}

type errorReceiver struct {
	err error
}

func (s *errorReceiver) errorTimes(err error) int {
	result := 0
	if s.err == nil {
		result = 1
	}

	if err != nil {
		s.err = err
	}

	return result
}
