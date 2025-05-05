package models

type User struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Password       string  `json:"password"`
	Income         float32 `json:"income"`
	Accountbalance float32 `json:"accountbalance"`
}
