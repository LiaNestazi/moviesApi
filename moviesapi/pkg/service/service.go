package service

import (
	moviesapi "github.com/LiaNestazi/moviesApi"
	"github.com/LiaNestazi/moviesApi/pkg/repository"
)

type Authorization interface {
	CreateUser(user moviesapi.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type MoviesList interface {
	GetMovies() ([]moviesapi.Movie, error)
	CreateMovie(movie moviesapi.Movie, actors_ids []int) (int, error)
	UpdateMovie(movieId int, input moviesapi.Movie) error
	DeleteMovie(movieId int) error
	//SearchForMovie(input string) ([]moviesapi.Movie, error)
}
type ActorsList interface {
	GetActors() ([]moviesapi.Actor, error)
	CreateActor(actor moviesapi.Actor) (int, error)
	UpdateActor(actorId int, input moviesapi.Actor) error
	DeleteActor(actorId int) error
	//SearchForActor(input string) ([]moviesapi.Actor, error)
}

type Service struct {
	Authorization
	MoviesList
	ActorsList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		MoviesList: NewMoviesListService(repos.MoviesList),
		ActorsList: NewActorsListService(repos.ActorsList),
	}
}