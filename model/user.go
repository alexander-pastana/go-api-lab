package model

type User struct {
	ID       int    `json:"id_user"`
	Name     string `json:"username"`
	Password string `json:"password"`
}
