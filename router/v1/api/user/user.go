package user

import (
	"github.com/Peterliang233/debate/errmsg"
	"github.com/Peterliang233/debate/model"
	"github.com/Peterliang233/debate/service/v1/api/user"
	"github.com/Peterliang233/debate/service/v1/api/user/validate"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)


//获取用户信息
func GetUser(c *gin.Context) {
	email, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errmsg.Error,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[errmsg.Error],
			},
		})
		return
	}

	StatusCode, code, u := user.GetUserInfo(email.(string))
	c.JSON(StatusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"data": u,
			"detail": errmsg.CodeMsg[code],
		},
	})
}

//更新用户信息
func UpdateUser(c *gin.Context) {
	var u model.User
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.ErrRequest,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[errmsg.ErrRequest],
			},
		})
	}

	StatusCode, code := user.UpdateUserInfo(&u)
	c.JSON(StatusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"detail": errmsg.CodeMsg[code],
			"data": u,
		},
	})
}

//上传文件接口
func UpdatePhoto(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.ErrRequest,
			"msg": map[string]interface{}{
				"status": errmsg.CodeMsg[errmsg.ErrRequest],
			},
		})
	} else {
		dist := path.Join("./", file.Filename)
		code := c.SaveUploadedFile(file, dist)
		if code != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": errmsg.Error,
				"msg": map[string]interface{}{
					"data":   "",
					"detail": "upload error",
				},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": errmsg.Success,
				"msg": map[string]interface{}{
					"data":   "",
					"detail": "upload success",
				},
			})
		}
	}
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

	StatusCode, code := user.UpdatePassword(&NewPassword)
	c.JSON(StatusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"detail": errmsg.Success,
		},
	})
}
