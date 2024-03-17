package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get(authorizationHeader)
	if header == "" {
		log.Fatalf("Empty auth header: %d", http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		log.Fatalf("Invalid auth heander: %d", http.StatusUnauthorized)
		return
	}

	if len(headerParts[1]) == 0 {
		log.Fatalf("Token is empty: %d", http.StatusUnauthorized)
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		log.Fatalf("Internal error occured: %d", http.StatusUnauthorized)
		return
	}
	ctx := r.Context()
	req := r.WithContext(context.WithValue(ctx, userCtx, userId))
	*r = *req
}

func getUserId(w http.ResponseWriter, r *http.Request) (int, error) {
	ctx := r.Context()
	id := ctx.Value(userCtx)
	if id == nil {
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
