package models

type User struct {
	BaseModel
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
