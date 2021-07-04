package main

import (
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

var CookieName = "api_cookie"

func main() {
	router := chi.NewRouter()
	router.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins:   []string{"https://frontend.test.local"},
				AllowedMethods:   []string{"GET", "POST"},
				AllowedHeaders:   []string{"Content-Type"},
				AllowCredentials: true,
			},
		),
	)
	router.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		http.SetCookie(rw, &http.Cookie{
			Name:   CookieName,
			Value:  "Secret",
			Domain: "test.local",
			// Domain:   "api.test.local",
			MaxAge:   60 * 10,
			HttpOnly: true,
			// SameSite Attribute is Strict
			SameSite: http.SameSiteStrictMode,
		})
	})
	router.Post("/name", func(rw http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(CookieName)
		if err != nil {
			log.Println(err)
			rw.WriteHeader(500)
			return
		}

		b, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Println(err)
			rw.WriteHeader(500)
			return
		}
		log.Printf("send body %s", string(b))
		log.Printf("api_cookie was sended: %s", cookie.Value)
	})

	http.ListenAndServe(":8080", router)
}
