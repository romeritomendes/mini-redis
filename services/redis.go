package services

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/romeritomendes/mini-redis/store"
)

type Client struct {
	db *store.Entity
}

func newClient(db *store.Entity) *Client {
	return &Client{
		db,
	}
}

func HandleClient(conn net.Conn, db *store.Entity) {
	defer conn.Close()

	client := newClient(db)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		resp := client.Execute(ln)
		fmt.Println(resp)
	}
}

func (cl *Client) Execute(commands string) string {
	parts := strings.Fields(strings.TrimSpace(commands))

	switch strings.ToUpper(parts[0]) {
	case "PING":
		return cl.db.PING()

	case "SET":
		return cl.db.SET(parts[1], parts[2])

	case "GET":
		return cl.db.GET(parts[1])

	default:
		return parts[0]
	}
}
