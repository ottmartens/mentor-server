package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ottmartens/mentor-server/controllers"
	"github.com/ottmartens/mentor-server/models"
	"github.com/ottmartens/mentor-server/utils"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.Use(models.JwtAuthentication)

	router.HandleFunc("/api/health", utils.HealthCheck).Methods("GET")

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/user/edit", controllers.EditUserProfile).Methods("POST")
	router.HandleFunc("/api/user/self", controllers.GetUserSelf).Methods("GET")
	router.HandleFunc("/api/user/image", controllers.GetUserImage).Methods("POST")
	router.HandleFunc("/api/user/{id}", controllers.GetUserProfile).Methods("GET")
	router.HandleFunc("/api/user/{id}", controllers.DeleteAccount).Methods("DELETE")

	router.HandleFunc("/api/groups", controllers.GetGroups).Methods("GET")
	router.HandleFunc("/api/groups/my-group", controllers.GetUsersGroup).Methods("GET")
	router.HandleFunc("/api/groups/{id}", controllers.GetGroupDetails).Methods("GET")
	router.HandleFunc("/api/groups/edit-group", controllers.EditGroupProfile).Methods("POST")

	router.HandleFunc("/api/available-mentors", controllers.GetAvailableMentors).Methods("GET")
	router.HandleFunc("/api/groups/request-creation", controllers.RequestGroupForming).Methods("POST")
	router.HandleFunc("/api/groups/accept-creation", controllers.HandleForming).Methods("POST")

	router.HandleFunc("/api/groups/join", controllers.RequestGroupJoining).Methods("POST")
	router.HandleFunc("/api/groups/accept-joining", controllers.HandleJoining).Methods("POST")

	router.HandleFunc("/api/template-activities", controllers.GetTemplateActivities).Methods("GET")

	router.HandleFunc("/api/activity", controllers.AddGroupActivity).Methods("POST")
	router.HandleFunc("/api/activity/image", controllers.UploadActivityImage).Methods("POST")
	router.HandleFunc("/api/activity/{id}", controllers.GetActivity).Methods("GET")

	router.HandleFunc("/api/global-settings", controllers.GetGlobalSettings).Methods("GET")

	// Admin routes
	router.HandleFunc("/api/user/verify", controllers.VerifyUser).Methods("POST")
	router.HandleFunc("/api/activity/verify", controllers.VerifyActivity).Methods("POST")
	router.HandleFunc("/api/template-activities", controllers.AddTemplateActivity).Methods("POST")
	router.HandleFunc("/api/all-users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/unverified-activities", controllers.GetUnverifiedActivities).Methods("GET")
	router.HandleFunc("/api/set-global-settings", controllers.SetGlobalSettings).Methods("POST")

	// Temporary dev routes
	router.HandleFunc("/api/group/create", controllers.CreateGroupDirectly).Methods("POST")

	// File server
	router.PathPrefix("/api/images/").Handler(http.StripPrefix("/api/images/", http.FileServer(http.Dir("./images/"))))

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe("0.0.0.0:8080", handlers.CORS(headersOk, originsOk, methodsOk)(router))

	if err != nil {
		fmt.Println(err)
	}
}
