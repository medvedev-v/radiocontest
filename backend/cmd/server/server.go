package server

import (
	"log"
	"net/http"

	mux "github.com/gorilla/mux"

	handler "github.com/medvedev-v/radiocontest/internal/handler"
	repository "github.com/medvedev-v/radiocontest/internal/repository"
	mysql "github.com/medvedev-v/radiocontest/pkg/mysql"
)

func StartAndServe() {
	db, err := mysql.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Роутинг
	r := mux.NewRouter()

	userRepo := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo)
	//r.HandleFunc("/echo", echoHandler.Echo).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	/**teammateRepo := repository.NewTeammateRepository(db)
	teammateHandler := handler.NewTeammateHandler(teammateRepo)
	r.HandleFunc("/teammates", teammateHandler.CreateUser).Methods("POST")
	r.HandleFunc("/teammates/{id}", teammateHandler.GetUser).Methods("GET")
	r.HandleFunc("/teammates/{id}", teammateHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/teammates/{id}", teammateHandler.DeleteUser).Methods("DELETE")**/

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
