package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"log"
	moviesapi "github.com/LiaNestazi/moviesApi"
)

type Actors []moviesapi.Actor

func (h *Handler) actorsActions(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//Получение всех актеров из БД
		h.getActors(w, r)
	}
	if r.Method == http.MethodPost {
		//Создание нового актера
		h.createActor(w, r)
	}
	if r.Method == http.MethodPut {
		//Изменение существующего актера
		h.updateActor(w, r)
	}
	if r.Method == http.MethodDelete {
		//Удаление существующего актера
		h.deleteActor(w, r)
	}

}

func (h *Handler) createActor(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	gender := r.FormValue("gender")
	birthday := r.FormValue("birthday")
	

	var input moviesapi.Actor
	input.Name = name
	input.Gender = gender
	input.Birthday = birthday

	id, err := h.services.ActorsList.CreateActor(input)
	if err != nil {
		log.Fatalf("Internal server error: %s", err.Error())
		return
	}
	log.Printf("Actor created successfully, id: %d", id)
}

func (h *Handler) getActors(w http.ResponseWriter, r *http.Request){
	items, err := h.services.ActorsList.GetActors()
	if err != nil {
		log.Fatalf("Internal server error: %s", err.Error())
		return
	}
	json.NewEncoder(w).Encode(items)
}

func (h *Handler) updateActor(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	actorId, err := strconv.Atoi(id)
	if err != nil{
		log.Fatalf("Error while parsing id: %s", err.Error())
		return
	}

	name := r.FormValue("name")
	gender := r.FormValue("gender")
	birthday := r.FormValue("birthday")

	var input moviesapi.Actor
	input.Name = name
	input.Gender = gender
	input.Birthday = birthday

	if err := h.services.ActorsList.UpdateActor(actorId, input); err != nil {
		log.Fatalf("Internal server error: %s", err.Error())
		return
	}
}

func (h *Handler) deleteActor(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	actorId, err := strconv.Atoi(id)
	if err != nil{
		log.Fatalf("Error while parsing id: %s", err.Error())
		return
	}
	err = h.services.ActorsList.DeleteActor(actorId)
	if err != nil {
		log.Fatalf("Internal server error: %s", err.Error())
		return
	}
}

func (h *Handler) searchForActor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Search for actor")
}
