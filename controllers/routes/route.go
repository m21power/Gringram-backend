package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	auth "github.com/m21power/GrinGram/Auth"
	"github.com/m21power/GrinGram/controllers/database"
	"github.com/m21power/GrinGram/controllers/handlers"
	"github.com/m21power/GrinGram/db"
	_ "github.com/m21power/GrinGram/docs"
	"github.com/m21power/GrinGram/usecases"
	httpSwagger "github.com/swaggo/http-swagger"
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
		log.Println("Can't connect to the database!")
		return
	}
	// Swagger route
	r.route.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Initialize dependencies
	userStore := database.UserNewStore(db)
	userUsecase := usecases.NewUserUsecase(userStore)
	userHandler := handlers.NewUserHandler(userUsecase)

	postStore := database.NewPostStore(db)
	postUsecase := usecases.NewPostRepository(postStore)
	postHandler := handlers.NewPostHandler(postUsecase)

	// User Routes
	userRoutes := r.route.PathPrefix("/api/v1").Subrouter()
	userRoutes.Handle("/signup", http.HandlerFunc(userHandler.CreateUser)).Methods("POST")
	userRoutes.Handle("/login", http.HandlerFunc(userHandler.Login)).Methods("POST")
	userRoutes.Handle("/users/me", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(userHandler.GetMe))).Methods("GET")
	userRoutes.Handle("/users/logout", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(userHandler.Logout))).Methods("POST")
	userRoutes.Handle("/users/{id}", auth.RoleMiddleware("admin")(http.HandlerFunc(userHandler.GetUserByID))).Methods("GET")
	userRoutes.Handle("/users/email/{email}", auth.RoleMiddleware("admin")(http.HandlerFunc(userHandler.GetUserByEmail))).Methods("GET")
	userRoutes.Handle("/users/username/{username}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(userHandler.GetUserByUsername))).Methods("GET")
	userRoutes.Handle("/users/delete/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(userHandler.DeleteUser))).Methods("DELETE")
	userRoutes.Handle("/users/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(userHandler.UpdateUser))).Methods("PUT")
	userRoutes.Handle("/users/image/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(userHandler.DeleteUserImage))).Methods("DELETE")

	// Post Routes
	postRoutes := r.route.PathPrefix("/api/v1").Subrouter()
	postRoutes.Handle("/posts", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.CreatePost))).Methods("POST")
	postRoutes.Handle("/posts/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.GetPostByID))).Methods("GET")
	postRoutes.Handle("/posts", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.GetPosts))).Methods("GET")
	postRoutes.Handle("/posts/user/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.GetPostsByUserID))).Methods("GET")
	postRoutes.Handle("/posts/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.DeletePost))).Methods("DELETE")
	postRoutes.Handle("/posts/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.UpdatePost))).Methods("PUT")

	// Comment Routes
	commentRoutes := r.route.PathPrefix("/api/v1").Subrouter()
	commentRoutes.Handle("/comments", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.CreateComment))).Methods("POST")
	commentRoutes.Handle("/comments/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.GetCommentByID))).Methods("GET")
	commentRoutes.Handle("/comments/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.UpdateComment))).Methods("PUT")
	commentRoutes.Handle("/comments/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.DeleteComment))).Methods("DELETE")

	// Like Routes
	likeRoutes := r.route.PathPrefix("/api/v1").Subrouter()
	likeRoutes.Handle("/likes", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.MakeLike))).Methods("POST")
	likeRoutes.Handle("/likes/{id}", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.GetLikers))).Methods("GET")

	// Interaction Routes
	interactionRoutes := r.route.PathPrefix("/api/v1").Subrouter()
	interactionRoutes.Handle("/posts/interactions/feed", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.GetFeed))).Methods("GET")
	interactionRoutes.Handle("/posts/interactions/view", auth.RoleMiddleware("admin", "user")(http.HandlerFunc(postHandler.ViewPost))).Methods("POST")
	interactionRoutes.Handle("/posts/interactions/waiting/{id}", auth.RoleMiddleware("admin")(http.HandlerFunc(postHandler.UpdateWaitingList))).Methods("PUT")
}

func (r *Router) Run(addr string, router *mux.Router) error {
	log.Println("Server running on port: ", addr)
	return http.ListenAndServe(addr, router)
}
