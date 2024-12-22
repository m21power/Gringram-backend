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
				utils.WriteError(w, err)
				return
			}
			token := cookie.Value
			_, userRole, err := GetUsernameAndRole(token)
			if err != nil {
				utils.WriteError(w, err)
				return
			}
			for _, role := range allowedRoles {
				if userRole == role {
					next.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, "Forbidden: Access is denied", http.StatusForbidden)
		})
	}
}
