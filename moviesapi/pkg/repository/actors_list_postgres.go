package repository

import (
	"fmt"
	"log"
	"strings"

	moviesapi "github.com/LiaNestazi/moviesApi"
	"github.com/jmoiron/sqlx"
)

type ActorsListPostgres struct {
	db *sqlx.DB
}

func NewActorsListPostgres(db *sqlx.DB) *ActorsListPostgres {
	return &ActorsListPostgres{db: db}
}

func (r *ActorsListPostgres) CreateActor(actor moviesapi.Actor) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var actorId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (name, gender, birthday) values ($1, $2, $3) RETURNING id", actorsListTable)

	row := tx.QueryRow(createItemQuery, actor.Name, actor.Gender, actor.Birthday)
	err = row.Scan(&actorId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return actorId, tx.Commit()
}

func (r *ActorsListPostgres) GetActors() ([]moviesapi.Actor, error) {
	var actors []moviesapi.Actor
	query := fmt.Sprintf(`SELECT ac.name, ac.gender, ac.birthday FROM %s ac LEFT JOIN %s am on am.actor_id=ac.id LEFT JOIN %s mv on am.movie_id=mv.id ORDER BY ac.id`,
		actorsListTable, moviesActorsTable, moviesListTable)
	if err := r.db.Select(&actors, query); err != nil {
		return nil, err
	}

	return actors, nil
}

func (r *ActorsListPostgres) DeleteActor(actorId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1",
		actorsListTable)
	_, err := r.db.Exec(query, actorId)

	return err
}

func (r *ActorsListPostgres) UpdateActor(actorId int, input moviesapi.Actor) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, input.Name)
		argId++
	}

	if input.Gender != "" {
		setValues = append(setValues, fmt.Sprintf("gender=$%d", argId))
		args = append(args, input.Gender)
		argId++
	}

	if input.Birthday != "" {
		setValues = append(setValues, fmt.Sprintf("birthday=$%d", argId))
		args = append(args, input.Birthday)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d",
		actorsListTable, setQuery, actorId)
	args = append(args, actorId)

	log.Printf("updateQuery: %s", query)
	log.Printf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
