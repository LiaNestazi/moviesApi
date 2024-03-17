package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

const (
	usersTable        = "users"
	moviesListTable   = "movies"
	actorsListTable   = "actors"
	moviesActorsTable = "actors_movies"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	//Таблица пользователей
	prep, err := db.Prepare(`
    CREATE TABLE IF NOT EXISTS users(
        id serial PRIMARY KEY,
        username VARCHAR(20) NOT NULL UNIQUE,
        password VARCHAR NOT NULL,
		role VARCHAR(5) NOT NULL
	);
    `)
	if err != nil {
		return nil, err
	}
	_, err = prep.Exec()
	if err != nil {
		return nil, err
	}

	//Таблица актеров
	prep, err = db.Prepare(`
	CREATE TABLE IF NOT EXISTS actors(
        id serial PRIMARY KEY,
        name VARCHAR(20) NOT NULL,
        gender VARCHAR(6) NOT NULL,
		birthday VARCHAR(10) NOT NULL
	);
	`)
	if err != nil {
		return nil, err
	}
	_, err = prep.Exec()
	if err != nil {
		return nil, err
	}

	//Таблица фильмов
	prep, err = db.Prepare(`
	CREATE TABLE IF NOT EXISTS movies(
		id serial PRIMARY KEY,
		name VARCHAR(150) NOT NULL,
        description VARCHAR(1000) NOT NULL,
		date VARCHAR(10) NOT NULL,
		rating NUMERIC(2,1) NOT NULL
	);
	`)
	if err != nil {
		return nil, err
	}
	_, err = prep.Exec()
	if err != nil {
		return nil, err
	}
	//Таблица соответствия актеров и фильмов
	prep, err = db.Prepare(`
	CREATE TABLE IF NOT EXISTS actors_movies(
		id serial PRIMARY KEY,
		actor_id INTEGER NOT NULL,
		movie_id INTEGER NOT NULL
	);
	`)
	if err != nil {
		return nil, err
	}
	_, err = prep.Exec()
	if err != nil {
		return nil, err
	}

	return db, nil
}
