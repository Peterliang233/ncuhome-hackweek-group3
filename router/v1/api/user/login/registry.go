package login

import (
	"github.com/Peterliang233/debate/dao"
	"github.com/Peterliang233/debate/errmsg"
	"github.com/Peterliang233/debate/model"
	Service "github.com/Peterliang233/debate/service/v1/api/user/login"
	"github.com/Peterliang233/debate/service/v1/api/user/validate"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//注册
func Registry(c *gin.Context) {
	var NewUser model.RegistryRequest
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
	StatusCode, Code := Service.CheckEmail(NewUser.Email)
	if Code != errmsg.Success {
		c.JSON(StatusCode, gin.H{
			"code": Code,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[Code],
			},
		})
		return
	}

	//等待验证邮箱，五分钟内有效
	if !Service.Validation(NewUser.Email, NewUser.Code) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.ErrEmailCode,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[errmsg.ErrEmailCode],
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
	//成功创建新用户，同时初始化用户得分
	if code == errmsg.Success {
		_, err := dao.Conn.Do("SET", u.Email + "score", "0")  //初始化每个用户的分数
		if err != nil {
			log.Fatal("初始化分数失败")
		}
	}
	c.JSON(StatusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"detail": errmsg.CodeMsg[code],
		},
	})
}


//获取邮箱验证码
type email struct{
<<<<<<< HEAD
	Email string `json:"email"`
=======
	Email string  `json:"email"`
>>>>>>> 98b1ce7affda920fbace65bf09d01e695060c99a
}
func GetEmailCode(c *gin.Context) {
	var code int
	var e email
	_ = c.ShouldBind(&e)
	//fmt.Printf(email)
	//发送邮箱验证码
	var emailCode string
	emailCode, code = Service.SendEmail(e.Email)
	if code != errmsg.Success {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": code,
			"msg": map[string]interface{}{
				"detail": "发送失败",
			},
		})
	}else{
		if !Service.SetRedis(e.Email, emailCode) {  //将生成的验证码保存在redis缓存里面
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": errmsg.ErrRedisCached,
				"msg": map[string]interface{}{
					"detail": errmsg.CodeMsg[errmsg.ErrRedisCached],
				},
			})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"code": errmsg.Success,
				"msg": map[string]interface{}{
					"email_code": emailCode,
					"detail": errmsg.CodeMsg[errmsg.Success],
				},
			})
		}
	}
}

