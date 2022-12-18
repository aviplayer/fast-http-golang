package dto

type User struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}
