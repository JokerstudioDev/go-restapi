package main

import (
	"time"
)

type Todo struct {
	ID        string `bson:"_id"`
	Name      string
	Completed bool
	Due       time.Time
}
