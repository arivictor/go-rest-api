package main

import (
	"github.com/arivictor/go-api/internal/comment"
	"github.com/arivictor/go-api/internal/database"
	"github.com/arivictor/go-api/internal/server"
	"github.com/arivictor/go-api/internal/transport"
	"github.com/gorilla/mux"
	"log"
)

// Responsible for instantiation of startup of application
func Run() error {
	log.Println("Connecting to database...")
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	log.Println("Migrating database...")
	err = db.MigrateDatabase()
	if err != nil {
		return err
	}

	// Provision router
	router := mux.NewRouter()

	// Sets up our middleware functions
	router.Use(transport.JSONMiddleware)

	// we also want to log every incoming request
	router.Use(transport.LoggingMiddleware)

	// We want to timeout all requests that take longer than 15 seconds
	router.Use(transport.TimeoutMiddleware)

	// Handles database interactions
	commentRepository := comment.NewRepository(db)

	// Handles business logic
	commentService := comment.NewService(commentRepository)

	// Handles HTTP requests
	commentHandler := comment.NewHandler(commentService)

	// Route request to handlers
	router.HandleFunc("/api/v1/comment/{id}", transport.Authorize(commentHandler.Delete)).Methods("DELETE")
	router.HandleFunc("/api/v1/comment/{id}", transport.Authorize(commentHandler.Update)).Methods("PATCH")
	router.HandleFunc("/api/v1/comment", transport.Authorize(commentHandler.Create)).Methods("POST")
	router.HandleFunc("/api/v1/comment/{id}", commentHandler.Get).Methods("GET")
	router.HandleFunc("/api/v1/comment", commentHandler.List).Methods("GET")

	// Provision HTTP server
	s := server.NewServer("8080", router)

	// Start Server
	err = s.Serve()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := Run()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
