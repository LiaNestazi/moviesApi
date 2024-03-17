package server

import (
	"encoding/json"
	"log"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	moviesapi "github.com/LiaNestazi/moviesApi"
)

type Movies []moviesapi.Movie

func (h *Handler) moviesAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//Получение всех фильмов из БД
		h.getMovies(w, r)
	}
	if r.Method == http.MethodPost {
		//Добавление нового фильма
		h.createMovie(w, r)
	}
	if r.Method == http.MethodPut {
		//Изменение существующего фильма
		h.updateMovie(w, r)
	}
	if r.Method == http.MethodDelete {
		//Удаление существующего фильма
		h.deleteMovie(w, r)
	}

}

func (h *Handler) createMovie(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	description := r.FormValue("description")
	rating := r.FormValue("rating")
	rating_float, err := strconv.ParseFloat(rating, 32)
	if err != nil{
		log.Fatalf("Error while parsing rating: %s", err.Error())
		return
	}

	date := r.FormValue("date")
	actors_ids := r.FormValue("actors")
	actorsList := strings.Split(actors_ids, ",")
	actorsListInt := []int{}


	for _,i := range actorsList{
		j, err := strconv.Atoi(i)
		if err != nil{
			log.Fatalf("Error while parsing actors ids: %s", err.Error())
			return
		}
		actorsListInt = append(actorsListInt, j)
	}


	var input moviesapi.Movie
	input.Name = name
	input.Description = description
	input.Rating = float32(rating_float)
	input.Date = date
	id, err := h.services.MoviesList.CreateMovie(input, actorsListInt)
	if err != nil {
		log.Fatalf("Internal server error: %s", err.Error())
		return
	}
	log.Printf("Movie created successfully, id: %d", id)

}
func (h *Handler) getMovies(w http.ResponseWriter, r *http.Request){
	items, err := h.services.MoviesList.GetMovies()
	if err != nil {
		log.Fatalf("Internal server error: %s", err.Error())
		return
	}
	json.NewEncoder(w).Encode(items)
}

func (h *Handler) updateMovie(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	movieId, err := strconv.Atoi(id)
	if err != nil{
		log.Fatalf("Error while parsing id: %s", err.Error())
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	rating := r.FormValue("rating")
	rating_float, err := strconv.ParseFloat(rating, 32)
	if err != nil{
		log.Fatalf("Error while parsing rating: %s", err.Error())
		return
	}

	date := r.FormValue("date")

	var input moviesapi.Movie
	input.Name = name
	input.Description = description
	input.Rating = float32(rating_float)
	input.Date = date

	if err := h.services.MoviesList.UpdateMovie(movieId, input); err != nil {
		log.Fatalf("Internal server error: %s", err.Error())
		return
	}
}

func (h *Handler) deleteMovie(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	movieId, err := strconv.Atoi(id)
	if err != nil{
		log.Fatalf("Error while parsing id: %s", err.Error())
		return
	}
	err = h.services.MoviesList.DeleteMovie(movieId)
	if err != nil {
		log.Fatalf("Internal server error: %s", err.Error())
		return
	}

}

func (h *Handler) searchForMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Search for movie")
}
