package main

import "time"

type Page struct {
	Title     string
	Body      string
	CreatedAt time.Time `json:"created_at"`
	EditedAT  time.Time `json:"edited_at"`
}
