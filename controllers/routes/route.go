package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	auth "github.com/m21power/GrinGram/Auth"
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

	r.route.Handle("/user/signup", http.HandlerFunc(userHandler.CreateUser)).Methods("POST")
	r.route.Handle("/user/login", http.HandlerFunc(userHandler.Login)).Methods("POST")
	r.route.Handle("/user/getme", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(userHandler.GetMe))).Methods("GET")
	r.route.Handle("/user/logout", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(userHandler.Logout))).Methods("POST")

	r.route.Handle("/user/{id}", auth.RoleMiddleware("admin")(http.HandlerFunc(userHandler.GetUserByID))).Methods("GET")
	r.route.Handle("/user/email/", auth.RoleMiddleware("admin")(http.HandlerFunc(userHandler.GetUserByEmail))).Methods("GET")
	r.route.Handle("/user/username/", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(userHandler.GetUserByUsername))).Methods("GET")

	//1
	r.route.Handle("/user/delete/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(userHandler.DeleteUser))).Methods("DELETE")
	r.route.Handle("/user/update/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(userHandler.UpdateUser))).Methods("PUT")
	r.route.Handle("/user/image/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(userHandler.DeleteUserImage))).Methods("DELETE")

	// post route
	postStore := database.NewPostStore(db)
	postUsecase := usecases.NewPostRepository(postStore)
	postHandler := handlers.NewPostHandler(postUsecase)

	r.route.Handle("/user/post", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.CreatePost))).Methods("POST")
	r.route.Handle("/user/post/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.GetPostByID))).Methods("GET")
	r.route.Handle("/user/posts/", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.GetPosts))).Methods("GET")
	r.route.Handle("/user/post/user/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.GetPostsByUserID))).Methods("GET")

	//2
	r.route.Handle("/user/post/delete/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.DeletePost))).Methods("DELETE")
	r.route.Handle("/user/post/update/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.UpdatePost))).Methods("PUT")

	//comment route
	r.route.Handle("/user/post/comment", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.CreateComment))).Methods("POST")
	r.route.Handle("/user/post/comment/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.GetCommentByID))).Methods("GET")
	//3
	r.route.Handle("/user/post/comment/update/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.UpdateComment))).Methods("PUT")
	r.route.Handle("/user/post/comment/delete/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.DeleteComment))).Methods("DELETE")
	//like route
	r.route.Handle("/user/post/like", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.MakeLike))).Methods("POST")
	r.route.Handle("/user/post/like/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.GetLikers))).Methods("GET")

	// interaction
	r.route.Handle("/user/post/feed/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.GetFeed))).Methods("GET")
	r.route.Handle("/user/post/view/", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.ViewPost))).Methods("POST")
	r.route.Handle("/user/post/waiting/{id}", auth.RoleMiddleware("admin")(http.HandlerFunc(postHandler.UpdateWaitingList))).Methods("PUT")
}

func (r *Router) Run(addr string, router *mux.Router) error {
	log.Println("Server running on port: ", addr)
	return http.ListenAndServe(addr, router)
}
