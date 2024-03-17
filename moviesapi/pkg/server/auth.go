package server

import (
	"fmt"
	"log"
	"net/http"
	moviesapi "github.com/LiaNestazi/moviesApi"
)

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		token, err := h.services.Authorization.GenerateToken(username, password)
		if err != nil {
			log.Fatalf("Internal server error: %s", err.Error())
			return
		}	

		log.Printf("Signed in successfully, token: %s", token)
		return

	} else {
		fmt.Fprintf(w, "This route accepts only POST requests")
	}
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		var input moviesapi.User
		input.Username = username
		input.Password = password

		id, err := h.services.Authorization.CreateUser(input)

		if err != nil {
			log.Fatalf("Internal server error: %s", err.Error())
			return
		}

		log.Printf("Signed up successfully: %d", id)
		return
	} else {
		fmt.Fprintf(w, "This route accepts only POST requests")
	}
}


// type AuthService struct{

// }

// func (s *AuthService) CreateUser(user moviesapi.User) (int, error) {
// 	user.Password = s.generatePasswordHash(user.Password)
// 	var id int
// 	query := fmt.Sprintf("INSERT INTO %s (username, password, role) values ($1, $2, $3) RETURNING id", usersTable)
	
// 	row := s.db.QueryRow(query, user.Username, user.Password, "user")
// 	if err := row.Scan(&id); err != nil {
// 		return 0, err
// 	}

// 	return id, nil
// }

// func (s *AuthService) generatePasswordHash(password string) string {
// 	hash := sha1.New()
// 	hash.Write([]byte(password))

// 	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
// }