package socket

import (
	"github.com/Peterliang233/debate/errmsg"
	"github.com/Peterliang233/debate/model"
	debate2 "github.com/Peterliang233/debate/service/v1/api/debate"
	"github.com/gin-gonic/gin"
	"net/http"
)


//1v1进行辩论
func OneToOneDebate(c *gin.Context) {
	var debate model.DebateRedis
	err := c.ShouldBind(&debate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.Error,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[errmsg.Error],
			},
		})
		return
	}

	StatusCode, code := debate2.CreateRecord(&debate)
	c.JSON(StatusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"detail": errmsg.CodeMsg[code],
		},
	})
	//content := strings.Split(debate.PositiveContent, " ")
}


//通过id获取辩论记录
func GetDebate(c *gin.Context) {
	id := c.Param("id")
	result, StatusCode, code := debate2.GetRedisHashRecord(id)
	c.JSON(StatusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"detail": errmsg.CodeMsg[code],
			"data": result,
		},
	})
}

//选择正方
func ChosePositive(c *gin.Context) {

}

//选择反方
func ChoseNegative(c *gin.Context) {

}