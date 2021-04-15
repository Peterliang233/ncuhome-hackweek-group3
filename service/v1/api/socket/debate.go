package socket

import (
	"github.com/Peterliang233/debate/dao"
	"github.com/Peterliang233/debate/errmsg"
	"github.com/Peterliang233/debate/model"
	"github.com/garyburd/redigo/redis"
	"net/http"
	"strconv"
	"time"
)


//将辩论内容记录在redis
func CreateRecord(debate * model.Debate) (int,int) {
	debateId, err := redis.Int(dao.Conn.Do("incr", "debateId"))
	if err != nil {
		return http.StatusInternalServerError, errmsg.Error
	}

	now := time.Now().Unix()
	_, err = dao.Conn.Do(
		"HMSET",
			"articleId" + strconv.Itoa(debateId),
			"title", debate.Title,
			"positive_content", debate.PositiveContent,
			"negative_content", debate.NegativeContent,
			"yid", debate.Yid,
			"nid", debate.Nid,
			"time", now,
	)
	if err != nil {
		return http.StatusInternalServerError, errmsg.Error
	}
	return http.StatusOK, errmsg.Success
}
