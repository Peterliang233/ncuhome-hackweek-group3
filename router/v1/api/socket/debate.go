package socket

import (
	"github.com/Peterliang233/debate/errmsg"
	"github.com/Peterliang233/debate/model"
	chat "github.com/Peterliang233/debate/service/v1/api/socket"
	"github.com/gin-gonic/gin"
	"net/http"
)


//1v1进行辩论
func OneToOneDebate(c *gin.Context) {
	var debate model.Debate
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
