package models

type User struct {
	Id        int    `json:"id" gorm:"index"`
	Username  string `json:"username"`
	About     string `json:"about"`
	Anonymous bool   `json:"anonymous"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}
