package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/iain-apw/wordle_game/ctx"
	"github.com/iain-apw/wordle_game/models"
)

func userFromToken(token string) *models.User {
	// FIXME: JWT, Oauth2 ...
	if token == "baz00ka" {
		return &models.User{
			ID:   "TestUserId",
			Name: "Daley Thompson",
		}
	}

	return nil
}

func authToken(r *http.Request) string {
	hdr := r.Header.Get("Authorization")
	return strings.TrimPrefix(hdr, "Bearer ")
}

func requireAuth(h http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		token := authToken(r)
		user := userFromToken(token)
		if user == nil {
			http.Error(w, "bad authentication", http.StatusUnauthorized)
			return
		}

		ctx, _ := ctx.AddUserToContext(user, r.Context())

		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func InitApiRoutes() (*mux.Router, error) {

	router := mux.NewRouter().StrictSlash(true)

	gameHandler, err := New()
	if err != nil {
		fmt.Printf("Error initialising API routes: %v\n", err)
		return nil, err
	}

	router.Methods("GET").Path("/games").Handler(http.HandlerFunc(gameHandler.GetAllGames))
	router.Methods("GET").Path("/games/{gameId}").Handler(http.HandlerFunc(gameHandler.GetGame))
	router.Methods("POST").Path("/games/create").Handler(requireAuth(http.HandlerFunc(gameHandler.CreateGame)))

	return router, nil
}

/*

type User struct {
	Login string
}

func userFromToken(token string) *User {
	// FIXME: JWT, Oauth2 ...
	if token == "baz00ka" {
		return &User{"joe"}
	}
	return nil
}

func authToken(r *http.Request) string {
	hdr := r.Header.Get("Authorization")
	return strings.TrimPrefix(hdr, "Bearer ")
}

func requireAuth(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		token := authToken(r)
		user := userFromToken(token)
		if user == nil {
			http.Error(w, "bad authentication", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
*/
