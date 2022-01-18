package migrations

type User struct {
	BaseModel
	Name string `json:"name"`
}
