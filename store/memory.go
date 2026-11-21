package store

import (
	"sync"
	"time"
)

type Entity struct {
	mu     sync.RWMutex
	store  map[string]string
	expiry map[string]time.Time
}

func NewEntity() *Entity {
	return &Entity{
		store:  make(map[string]string),
		expiry: make(map[string]time.Time),
	}
}

func (db *Entity) PING() string {
	return "PONG\r\n"
}

func (db *Entity) SET(key, value string) string {
	db.store[key] = value

	return "OK\r\n"
}

func (db *Entity) GET(key string) string {
	_, ok := db.store[key]
	if ok {
		return db.store[key]
	}

	return "NOK\r\n"
}

func (db *Entity) DEL(key string) string {
	_, ok := db.store[key]
	if ok {
		delete(db.store, key)
		return "OK\r\n"
	}

	return "NOK\r\n"
}
