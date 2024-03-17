package service

import(
	moviesapi "github.com/LiaNestazi/moviesApi"
	"github.com/LiaNestazi/moviesApi/pkg/repository"
)

type MoviesListService struct{
	repo repository.MoviesList
}

func NewMoviesListService(repo repository.MoviesList) *MoviesListService {
	return &MoviesListService{repo: repo}
}

func (s *MoviesListService) CreateMovie(movie moviesapi.Movie, actors_ids []int) (int, error) {
	return s.repo.CreateMovie(movie, actors_ids)
}

func (s *MoviesListService) GetMovies() ([]moviesapi.Movie, error) {
	return s.repo.GetMovies()
}

func (s *MoviesListService) DeleteMovie(movieId int) error {
	return s.repo.DeleteMovie(movieId)
}

func (s *MoviesListService) UpdateMovie(movieId int, input moviesapi.Movie) error {
	return s.repo.UpdateMovie(movieId, input)
}