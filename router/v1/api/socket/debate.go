package socket

import (
	"github.com/Peterliang233/debate/dao"
	"github.com/Peterliang233/debate/errmsg"
	"github.com/Peterliang233/debate/model"
	chat "github.com/Peterliang233/debate/service/v1/api/socket"
	"github.com/garyburd/redigo/redis"
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

	StatusCode, code := chat.CreateRecord(&debate)
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
	result, err := redis.Values(dao.Conn.Do("HGETALL", id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errmsg.Error,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[errmsg.Error],
			},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": errmsg.Success,
		"msg": map[string]interface{}{
			"detail": errmsg.CodeMsg[errmsg.Success],
			"data": result,
		},
	})
}