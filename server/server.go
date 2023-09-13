package server

import (
	repo "github.com/muhammadarash1997/movies/repository"
	svc "github.com/muhammadarash1997/movies/service"
)

type Movies struct {
	repo repo.Repo
}

func NewMoviesService(repo repo.Repo) svc.MoviesService {
	return &Movies{repo: repo}
}
