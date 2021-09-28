package database

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/yona3/go-auth-sample/ent"
)

var client *ent.Client

func Init() {
	c, err := ent.Open("postgres", "host=localhost port=5433 user=root dbname=go-auth-sample password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")
	client = c
}

func GetClient() *ent.Client {
	return client
}

func CloseClient() error {
	log.Println("Closing client")
	return client.Close()
}
