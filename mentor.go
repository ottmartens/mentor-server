package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ottmartens/mentor-server/app"
	"github.com/ottmartens/mentor-server/controllers"
	"github.com/ottmartens/mentor-server/utils"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)

	router.HandleFunc("/api/health", utils.HealthCheck).Methods("GET")
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe("localhost:8080", router)

	if err != nil {
		fmt.Println(err)
	}
}
