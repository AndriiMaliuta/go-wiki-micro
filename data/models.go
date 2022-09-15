package data

import "time"

type Page struct {
	Title     string
	Body      string
	CreatedAt time.Time `json:"created_at"`
	EditedAT  time.Time `json:"edited_at"`
}

type User struct {
	Name     string
	Key      string
	FullName string `json:"full_name"`
}

type Version struct {
	Number    int16
	CreatedAt time.Time `json:"created_at"`
	By        User
}

type Metadata struct {
	Tags string
}

type History struct {
	Metadata Metadata
	Version  Version
}

type Blog struct {
	Title     string
	Body      string
	History   History
	CreatedAt time.Time `json:"created_at"`
	EditedAT  time.Time `json:"edited_at"`
}

type Space struct {
	Title     string
	Key       string
	Category  string
	CreatedAt time.Time `json:"created_at"`
	EditedAT  time.Time `json:"edited_at"`
}
