package auth

import (
	"net/http"

	"github.com/m21power/GrinGram/utils"
)

// Middleware for role-based access
func RoleMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("token")
			if err != nil {
				utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
				return
			}
			token := cookie.Value
			Token, err := GetTokenValues(token)
			if err != nil {
				utils.WriteJSON(w, http.StatusBadRequest, utils.ApiResponse{Message: err.Error(), Success: false})
				return
			}
			for _, role := range allowedRoles {
				if Token.Role == role {
					next.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, "Forbidden: Access is denied", http.StatusForbidden)
		})
	}
}
