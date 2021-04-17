package login
//
//import (
//	"github.com/Peterliang233/debate/dao"
//	"github.com/Peterliang233/debate/errmsg"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//
////登出
//func Logout(c *gin.Context) {
//	email := c.Query("email")
//	//删除redis里面的token
//	_, err := dao.Conn.Do("DEL", email + "token")
//	StatusCode, code := http.StatusOK, errmsg.Success
//	if err != nil {
//		StatusCode = http.StatusInternalServerError
//		code = errmsg.Error
//	}
//	c.JSON(StatusCode, gin.H{
//		"code": code,
//		"msg": map[string]interface{}{
//			"detail": errmsg.CodeMsg[code],
//		},
//	})
//}
//
