package repository

import (
	"fmt"
	"log"
	"strings"

	moviesapi "github.com/LiaNestazi/moviesApi"
	"github.com/jmoiron/sqlx"
)

type MoviesListPostgres struct {
	db *sqlx.DB
}

func NewMoviesListPostgres(db *sqlx.DB) *MoviesListPostgres {
	return &MoviesListPostgres{db: db}
}

func (r *MoviesListPostgres) CreateMovie(movie moviesapi.Movie, actors_ids []int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var movie_id int
	createQuery := fmt.Sprintf("INSERT INTO %s (name, description, rating, date) VALUES ($1, $2, $3, $4) RETURNING id", moviesListTable)
	row := tx.QueryRow(createQuery, movie.Name, movie.Description, movie.Rating, movie.Date)
	if err := row.Scan(&movie_id); err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, id := range actors_ids {
		createActorsListQuery := fmt.Sprintf("INSERT INTO %s (actor_id, movie_id) VALUES ($1, $2)", moviesActorsTable)
		_, err = tx.Exec(createActorsListQuery, id, movie_id)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return movie_id, tx.Commit()
}

func (r *MoviesListPostgres) GetMovies() ([]moviesapi.Movie, error) {
	var movies []moviesapi.Movie

	query := fmt.Sprintf("SELECT name, description, rating, date FROM %s",
		moviesListTable)
	err := r.db.Select(&movies, query)

	return movies, err
}

func (r *MoviesListPostgres) DeleteMovie(movieId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1",
		moviesListTable)
	_, err := r.db.Exec(query, movieId)

	return err
}

func (r *MoviesListPostgres) UpdateMovie(movieId int, input moviesapi.Movie) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, input.Name)
		argId++
	}

	if input.Description != "" {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, input.Description)
		argId++
	}

	if input.Rating != 0 {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, input.Rating)
		argId++
	}

	if input.Date != "" {
		setValues = append(setValues, fmt.Sprintf("date=$%d", argId))
		args = append(args, input.Date)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d",
		moviesListTable, setQuery, movieId)
	args = append(args, movieId)

	log.Printf("updateQuery: %s", query)
	log.Printf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
