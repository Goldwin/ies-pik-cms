package dto

type Role struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Scopes []string `json:"scopes"`
}
