package login

import (
	//"context"
	"github.com/Peterliang233/debate/errmsg"
	"github.com/Peterliang233/debate/middlerware"
	"github.com/Peterliang233/debate/model"
	"github.com/Peterliang233/debate/router/v1/api/user/validate"
	Service "github.com/Peterliang233/debate/service/v1/api/user/login"
	"github.com/gin-gonic/gin"
	"net/http"
)


//登录
func Login(c *gin.Context) {
	var login model.Login
	err := c.ShouldBind(&login)
	msg, code := validate.Validate(&login)
	if code != errmsg.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg": msg,
		})
		return
	}
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
	token, code = middlerware.GenerateToken(login.Email)
	c.JSON(StatusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"token": token,
			"detail": errmsg.CodeMsg[code],
		},
	})
}


//更新用户密码
func UpdatePassword(c *gin.Context) {
	var NewPassword model.UpdateNewPassword
	err := c.ShouldBind(&NewPassword)
	msg, code := validate.Validate(&NewPassword)
	if code != errmsg.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg": msg,
		})
		return
	}
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