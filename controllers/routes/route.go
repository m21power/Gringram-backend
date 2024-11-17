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

	r.route.HandleFunc("/user", userHandler.CreateUser).Methods("POST")
	r.route.HandleFunc("/user/{id}", userHandler.GetUserByID).Methods("GET")
	r.route.HandleFunc("/user/email/", userHandler.GetUserByEmail).Methods("GET")
	r.route.HandleFunc("/user/username/", userHandler.GetUserByUsername).Methods("GET")
	r.route.HandleFunc("/user/{id}", userHandler.DeleteUser).Methods("DELETE")
	r.route.HandleFunc("/user/update/{id}", userHandler.UpdateUser).Methods("PUT")
	r.route.HandleFunc("/user/image/{id}", userHandler.DeleteUserImage).Methods("DELETE")

	// post route
	postStore := database.NewPostStore(db)
	postUsecase := usecases.NewPostRepository(postStore)
	postHandler := handlers.NewPostHandler(postUsecase)

	r.route.HandleFunc("/user/post", postHandler.CreatePost).Methods("POST")
	r.route.HandleFunc("/user/post/{id}", postHandler.GetPostByID).Methods("GET")
	r.route.HandleFunc("/user/post/user/{id}", postHandler.GetPostsByUserID).Methods("GET")
	r.route.HandleFunc("/user/post/{id}", postHandler.DeletePost).Methods("DELETE")
	r.route.HandleFunc("/user/post/update/{id}", postHandler.UpdatePost).Methods("PUT")

	// POST IMAGE ROUTE
	r.route.HandleFunc("/user/post/image", postHandler.CreatePostImage).Methods("POST")
	r.route.HandleFunc("/user/post/image/update/{id}", postHandler.UpdatePostImage).Methods("PUT")
	r.route.HandleFunc("/user/post/image/{id}", postHandler.DeletePostImage).Methods("DELETE")
	r.route.HandleFunc("/user/post/image/{id}", postHandler.GetPostImageByID).Methods("GET")

}

func (r *Router) Run(addr string, ru *mux.Router) error {
	log.Println("Server running on port: ", addr)
	return http.ListenAndServe(addr, ru)
}
