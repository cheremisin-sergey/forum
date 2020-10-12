package api

import "github.com/gin-gonic/gin"

type CreateForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

func CreateUser(context *gin.Context) {

}
