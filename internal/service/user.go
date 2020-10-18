package service

import (
	//"github.com/EDDYCJY/go-gin-example/models"
)

type User struct {
	Id        int    `json:"id" gorm:"index"`
	Username  string `json:"username"`
	About     string `json:"about"`
	Anonymous bool   `json:"anonymous"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

func (u *User) Create() error {
	//user := map[string]interface{}{
	//	"id":        u.Id,
	//	"username":  u.Username,
	//	"about":     u.About,
	//	"anonymous": u.Anonymous,
	//	"name":      u.Name,
	//	"email":     u.Email,
	//}
	//
	//
	//if err := models.AddArticle(article); err != nil {
	//	return err
	//}

	return nil
}
