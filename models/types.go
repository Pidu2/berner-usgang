package models

type Event struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Artists string `json:"artists"`
	Genre   string `json:"genre"`
	IsImage bool   `json:"isimage"`
}
