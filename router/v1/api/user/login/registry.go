package login

import (
	"github.com/Peterliang233/debate/errmsg"
	"github.com/Peterliang233/debate/model"
	"github.com/Peterliang233/debate/router/v1/api/user/validate"
	Service "github.com/Peterliang233/debate/service/v1/api/user/login"
	"github.com/gin-gonic/gin"
	"net/http"
)

//注册
func Registry(c *gin.Context) {
	var NewUser  model.User
	err := c.ShouldBind(&NewUser)
	//验证
	msg, code := validate.Validate(&NewUser)
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
	//检查用户名和电话是否存在
	StatusCode, code := Service.CheckUser(NewUser.Username, NewUser.Password)
	if code != errmsg.Success {
		c.JSON(StatusCode, gin.H{
			"code": code,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[code],
			},
		})
		return
	}
	//发送短信验证码


	StatusCode, code = Service.CreateUser(&NewUser)
	c.JSON(StatusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"detail": errmsg.CodeMsg[code],
		},
	})
}

