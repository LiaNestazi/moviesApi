package server

type Actor struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}

type Movie struct {
	Id          int     `json:"-"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Rating      float32 `json:"rating"`
}

type ActorsMovies struct {
	Id      int
	ActorId int
	MovieId int
}
