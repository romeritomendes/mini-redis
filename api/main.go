package main

import (
	"log"

	"github.com/romeritomendes/mini-redis/server"
)

func main() {
	if err := server.NewServer(); err != nil {
		log.Fatalln(err)
	}
}
