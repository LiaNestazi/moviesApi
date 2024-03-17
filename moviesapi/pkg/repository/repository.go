package repository

import (
	"github.com/jmoiron/sqlx"
	moviesapi "github.com/LiaNestazi/moviesApi"
)

type Authorization interface {
	CreateUser(user moviesapi.User) (int, error)
	GetUser(username, password string) (moviesapi.User, error)
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

type Repository struct {
	Authorization
	MoviesList
	ActorsList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		MoviesList: NewMoviesListPostgres(db),
		ActorsList: NewActorsListPostgres(db),
	}
}