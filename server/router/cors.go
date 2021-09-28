package router

import "github.com/rs/cors"

func NewCors() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET"},
	})

	return c
}
