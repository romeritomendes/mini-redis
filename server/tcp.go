package server

import (
	"fmt"
	"net"

	"github.com/romeritomendes/mini-redis/services"
	"github.com/romeritomendes/mini-redis/store"
)

func NewServer() error {
	conn, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Println("Server running and Listen on Port 6379")

	db := store.NewEntity()
	for {
		client, err := conn.Accept()
		if err != nil {
			return err
		}

		fmt.Println("- New connection Accepted!")
		go services.HandleClient(client, db)
	}
}
