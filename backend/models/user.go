package models

type UserStruct struct {
	Name string `json:"name"`
	Certificate string `json:"certificate"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Id string `json:"id"`
	Balance int `json:"balance"`
}
