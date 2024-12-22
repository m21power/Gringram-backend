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
	r.route.HandleFunc("/user/login", userHandler.Login).Methods("POST")
	r.route.HandleFunc("/user/logout", userHandler.Logout).Methods("POST")
	r.route.HandleFunc("/user/getme", userHandler.GetMe).Methods("GET")

	r.route.HandleFunc("/user/{id}", userHandler.GetUserByID).Methods("GET")
	r.route.HandleFunc("/user/email/", userHandler.GetUserByEmail).Methods("GET")
	r.route.HandleFunc("/user/username/", userHandler.GetUserByUsername).Methods("GET")
	r.route.HandleFunc("/user/delete/{id}", userHandler.DeleteUser).Methods("DELETE")
	r.route.HandleFunc("/user/update/{id}", userHandler.UpdateUser).Methods("PUT")
	r.route.HandleFunc("/user/image/{id}", userHandler.DeleteUserImage).Methods("DELETE")

	// post route
	postStore := database.NewPostStore(db)
	postUsecase := usecases.NewPostRepository(postStore)
	postHandler := handlers.NewPostHandler(postUsecase)

	r.route.HandleFunc("/user/post", postHandler.CreatePost).Methods("POST")
	r.route.HandleFunc("/user/post/{id}", postHandler.GetPostByID).Methods("GET")
	r.route.HandleFunc("/user/posts/", postHandler.GetPosts).Methods("GET")
	r.route.HandleFunc("/user/post/user/{id}", postHandler.GetPostsByUserID).Methods("GET")
	r.route.HandleFunc("/user/post/delete/{id}", postHandler.DeletePost).Methods("DELETE")
	r.route.HandleFunc("/user/post/update/{id}", postHandler.UpdatePost).Methods("PUT")
	// Secure routes with roleMiddleware

	// r.route.Handle("/user", roleMiddleware("public")(http.HandlerFunc(userHandler.CreateUser))).Methods("POST")
	// r.route.Handle("/user/{id}", roleMiddleware("user")(http.HandlerFunc(userHandler.GetUserByID))).Methods("GET")
	// r.route.Handle("/user/email/", roleMiddleware("admin")(http.HandlerFunc(userHandler.GetUserByEmail))).Methods("GET")
	// r.route.Handle("/user/username/", roleMiddleware("user")(http.HandlerFunc(userHandler.GetUserByUsername))).Methods("GET")
	// r.route.Handle("/user/delete/{id}", roleMiddleware("admin")(http.HandlerFunc(userHandler.DeleteUser))).Methods("DELETE")
	// r.route.Handle("/user/update/{id}", roleMiddleware("user")(http.HandlerFunc(userHandler.UpdateUser))).Methods("PUT")
	// r.route.Handle("/user/image/{id}", roleMiddleware("user")(http.HandlerFunc(userHandler.DeleteUserImage))).Methods("DELETE")

	//comment route
	r.route.HandleFunc("/user/post/comment", postHandler.CreateComment).Methods("POST")
	r.route.HandleFunc("/user/post/comment/update/{id}", postHandler.UpdateComment).Methods("PUT")
	r.route.HandleFunc("/user/post/comment/delete/{id}", postHandler.DeleteComment).Methods("DELETE")
	r.route.HandleFunc("/user/post/comment/{id}", postHandler.GetCommentByID).Methods("GET")

	//like route
	r.route.HandleFunc("/user/post/like", postHandler.MakeLike).Methods("POST")
	r.route.HandleFunc("/user/post/like/{id}", postHandler.GetLikers).Methods("GET")

	// interaction
	r.route.HandleFunc("/user/post/feed/{id}", postHandler.GetFeed).Methods("GET")
	r.route.HandleFunc("/user/post/view/", postHandler.ViewPost).Methods("POST")
	r.route.HandleFunc("/user/post/waiting/{id}", postHandler.UpdateWaitingList).Methods("PUT")
}

func (r *Router) Run(addr string, ru *mux.Router) error {
	log.Println("Server running on port: ", addr)
	return http.ListenAndServe(addr, ru)
}
