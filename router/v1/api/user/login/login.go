package login

import (
	//"context"
	"github.com/Peterliang233/debate/errmsg"
	"github.com/Peterliang233/debate/middlerware"
	"github.com/Peterliang233/debate/model"
	Service "github.com/Peterliang233/debate/service/v1/api/user/login"
	"github.com/gin-gonic/gin"
	//proto "github.com/Peterliang233/debate/service/v1/api/user/login/proto"
	"net/http"
)


//type UserService struct {
//	User *login.User
//}
//
//func (u *UserService) Register(ctx context.Context, req *proto.RegisterRequest, resq *proto.Response) error {
//
//}
//
//
//func (u *UserService) Login(ctx context.Context, req *proto.LoginRequest, resq *proto.Response) error {
//
//}
//
//func (u *UserService) UpdatePassword(ctx context.Context, req *proto.UpdatePasswordRequest, resq *proto.Response) error {
//
//}


//登录
func Login(c *gin.Context) {
	var login model.Login
	err := c.ShouldBind(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.ErrRequest,
			"detail": errmsg.CodeMsg[errmsg.ErrRequest],
		})
		return
	}
	StatusCode, code := Service.CheckLogin(&login)
	if code != errmsg.Success {
		c.JSON(StatusCode, gin.H{
			"code": code,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[code],
			},
		})
		return
	}
	var token string
	//获取token
	token, code = middlerware.GenerateToken(login.Phone)
	c.JSON(StatusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"token": token,
			"detail": errmsg.CodeMsg[code],
		},
	})
}

type UpdateNewPassword struct {
	Phone string `json:"phone"`
	OldPassword string 	`json:"old_password"`
	NewPassword string 	`json:"new_password"`
	CheckNewPassword string  `json:"check_new_password"`
}

func UpdatePassword(c *gin.Context) {
	var NewPassword UpdateNewPassword
	err := c.ShouldBind(&NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.ErrRequest,
			"detail": errmsg.CodeMsg[errmsg.ErrRequest],
		})
		return
	}

	StatusCode, code := Service.UpdatePassword(&NewPassword)
	c.JSON(StatusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"detail": errmsg.Success,
		},
	})
}