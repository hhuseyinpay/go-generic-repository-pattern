package models

type User struct {
	BaseModel
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Detail  Detail `json:"detail"`
}

type Detail struct {
	Phone string `json:"phone"`
	Type  int    `json:"type"`
}
