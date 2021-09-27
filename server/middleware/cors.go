package middleware

import "github.com/rs/cors"

func CORS() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET"},
	})

	return c
}
