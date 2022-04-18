package models

type User struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}
