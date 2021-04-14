package login

import (
	"fmt"
	"github.com/Peterliang233/debate/errmsg"
	"github.com/Peterliang233/debate/model"
	"github.com/Peterliang233/debate/router/v1/api/user/validate"
	Service "github.com/Peterliang233/debate/service/v1/api/user/login"
	"github.com/gin-gonic/gin"
	"net/http"
)

//注册
func Registry(c *gin.Context) {
	var NewUser model.RegistryQuest
	err := c.ShouldBind(&NewUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.ErrRequest,
			"detail": errmsg.CodeMsg[errmsg.ErrRequest],
		})
		return
	}
	//数据验证
	msg, code := validate.Validate(&NewUser)
	if code != errmsg.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg": msg,
		})
		return
	}
	//检查邮箱是否被用
	StatusCode, code := Service.CheckEmail(NewUser.Email)
	if code != errmsg.Success {
		c.JSON(StatusCode, gin.H{
			"code": code,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[code],
			},
		})
		return
	}

	//等待验证邮箱，五分钟内有效
	if !Service.Validation(NewUser.Code) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.ErrUserEmailUsed,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[errmsg.ErrUserEmailUsed],
			},
		})
		return
	}


	//验证成功
	u := &model.User{
		Email: NewUser.Email,
		Password: NewUser.Password,
	}
	StatusCode, code = Service.CreateUser(u)
	c.JSON(StatusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"detail": errmsg.CodeMsg[code],
		},
	})
}


//获取邮箱验证码
func GetEmailCode(c *gin.Context) {
	var code int
	email := c.Query("email")
	fmt.Printf(email)
	//发送邮箱验证码
	var emailCode string
	emailCode, code = Service.SendEmail(email)
	if code != errmsg.Success {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": code,
			"msg": map[string]interface{}{
				"detail": "发送失败",
			},
		})
	}else{
		if !Service.SetRedis(emailCode) {  //将生成的验证码保存在redis缓存里面
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "error",
			})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"msg": emailCode,
			})
		}
	}
}

