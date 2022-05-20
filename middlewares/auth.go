package middlewares

import (
	"context"
	"goapi/service"
	"log"
	"net/http"
)

type authString string
type custom_error string

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]
		// log.Printf(auth)
		validate, err := service.JwtValidate(context.Background(), auth)
		if err != nil || !validate.Valid {
			var custom_error = "{\"errors\": [{\"message\": \"Access token expired\"}],\"data\":null}"
			w.Write([]byte(custom_error))
			return
		}

		customClaim, _ := validate.Claims.(*service.JwtCustomClaim)
		log.Print(customClaim)
		ctx := context.WithValue(r.Context(), authString("auth"), customClaim)
		tokenStr := ctx.Value(authString("auth"))
		log.Println("Token get ....")
		log.Print(tokenStr)
		log.Print(customClaim.ID)
		log.Print(customClaim.Id)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func CtxValue(ctx context.Context) *service.JwtCustomClaim {
	raw, _ := ctx.Value(authString("auth")).(*service.JwtCustomClaim)
	return raw
}
