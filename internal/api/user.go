package api

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/cheremisin-sergey/forum/internal/constants"
	context2 "github.com/cheremisin-sergey/forum/internal/context"
	"github.com/cheremisin-sergey/forum/pkg/app"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username  string `json:"username" valid:"Required;MaxSize(254)"`
	About     string `json:"about" valid:"Required"`
	Anonymous bool   `json:"anonymous"`
	Name      string `json:"name" valid:"Required;MaxSize(254)"`
	Email     string `json:"email" valid:"Required;MaxSize(254)"`
}

type CreateUserResponse struct {
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	Response interface{} `json:"response"`
}

func CreateUser(c *gin.Context) {
	appContext := c.MustGet(constants.AppContext).(*context2.AppContext)

	var (
		request CreateUserRequest
	)

	httpCode, errCode := app.BindAndValid(c, &request)
	if errCode != e.SUCCESS {
		c.JSON(httpCode, CreateUserResponse{
			Code:     errCode,
			Msg:      e.GetMsg(errCode),
			Response: nil,
		})
		return
	}

	query := fmt.Sprintf(
		`INSERT INTO users(username,about,anonymous,name,email) VALUES('%s','%s',%t,'%s','%s') returning id;`,
		request.Username,
		request.About,
		request.Anonymous,
		request.Name,
		request.Email,
	)

	fmt.Println(query)


	var lastInsertId int64
	err := appContext.DB.QueryRow(query).Scan(&lastInsertId)
	if err != nil {
		panic(err)
	}

	fmt.Println(lastInsertId)

	c.JSON(httpCode, CreateUserResponse{
		Code:     200,
		Msg:      "Success" ,
		Response: nil,
	})



}

func GetUser(context *gin.Context) {

}
