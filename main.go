package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	//loading .env (environment variables)
	errgodot := godotenv.Load(".env")
	if errgodot != nil {
		log.Fatal("Error loading .env file")
	}
	//retrieving PORT variable data
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}
	fmt.Println("Port:", portString)

	//initiating router
	router := chi.NewRouter()

	//initiating cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//using seperate handler
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	// v1Router.HandleFunc("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	router.Mount("/v1", v1Router)

	//initiate server properties
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server starting on port %v\n", portString)
	// router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello, World!")
	// 	log.Println("Request for GET executed")

	// })
	// router.Post("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello, World!")
	// 	log.Println("Request for POST executed")
	// 	w.WriteHeader(500)
	// })
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server %v\n", err)
	}
}
