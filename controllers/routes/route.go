package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/m21power/GrinGram/controllers/database"
	"github.com/m21power/GrinGram/controllers/handlers"
	"github.com/m21power/GrinGram/db"
	"github.com/m21power/GrinGram/usecases"
)

type Router struct {
	route *mux.Router
}

func NewRouter(r *mux.Router) *Router {
	return &Router{route: r}
}

func (r *Router) RegisterRoute() {
	db, err := db.ConnectDB()
	if err != nil {
		log.Println("Can't connect to the databse!")
		return
	}
	userStore := database.UserNewStore(db)
	userUsecase := usecases.NewUserUsecase(userStore)
	userHandler := handlers.NewUserHandler(userUsecase)
	profileStore := database.NewProfileStore(db)
	profileUsecase := usecases.NewProfileUsecase(profileStore)
	profileHandler := handlers.NewProfileHandler(profileUsecase)

	r.route.HandleFunc("/user", userHandler.CreateUser).Methods("POST")
	r.route.HandleFunc("/user/{id}", userHandler.GetUserByID).Methods("GET")
	r.route.HandleFunc("/email", userHandler.GetUserByEmail).Methods("GET")
	r.route.HandleFunc("/username", userHandler.GetUserByUsername).Methods("GET")
	r.route.HandleFunc("/user/{id}", userHandler.DeleteUser).Methods("DELETE")
	r.route.HandleFunc("/user/{id}", userHandler.UpdateUser).Methods("PUT")

	r.route.HandleFunc("/user/image/{id}", profileHandler.CreateProfile).Methods("POST")
	r.route.HandleFunc("/user/image/{id}", profileHandler.GetProfileByID).Methods("GET")
	r.route.HandleFunc("/user/image", profileHandler.UpdateProfile).Methods("PUT")
	r.route.HandleFunc("/user/image/{id}", profileHandler.DeleteProfile).Methods("DELETE")

}

func (r *Router) Run(addr string, ru *mux.Router) error {
	log.Println("Server running on port: ", addr)
	return http.ListenAndServe(addr, ru)
}
