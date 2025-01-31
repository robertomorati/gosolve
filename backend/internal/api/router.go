package api

import (
	"net/http"

	_ "gosolve/backend/docs"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter(handler *Handler) *mux.Router {
	router := mux.NewRouter()

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	router.Handle("/search/{value}", http.HandlerFunc(handler.SearchHandler)).Methods("GET")

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"alive"}`))
	}).Methods("GET")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.Use(cors)
	return router
}
