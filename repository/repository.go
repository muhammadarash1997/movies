package repository

import (
	"io"

	"github.com/go-kit/kit/log/level"
	_interface "github.com/muhammadarash1997/movies/repository/interface"
	"github.com/muhammadarash1997/movies/repository/mysql"
	shv "github.com/muhammadarash1997/movies/utils/sharevar"
)

type DBConf struct {
	URL, Schema, User, Password string
}

type RepoConfigs struct {
	DBConf
}

type Repo struct {
	io.Closer
	DBReadWriter _interface.ReadWriter
}

func NewMoviesRepo(rc RepoConfigs) (*Repo, error) {
	// Database
	readWriter, err := mysql.NewDBReadWriter(rc.DBConf.URL, rc.DBConf.Schema, rc.DBConf.User, rc.DBConf.Password)
	if err != nil {
		return nil, err
	}

	return &Repo{
		DBReadWriter: readWriter,
	}, nil
}

func (r *Repo) Close() {
	level.Info(shv.Logger).Log("Closing Client")
	r.DBReadWriter.Close()
}
