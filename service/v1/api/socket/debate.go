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
func CreateRecord(debate * model.DebateRedis) (int,int) {
	debateId, err := redis.Int(dao.Conn.Do("incr", "debateId"))  //记录下辩论的总场次
	if err != nil {
		return http.StatusInternalServerError, errmsg.Error
	}

	deb := model.DebateMysql{
		Id: debate.Id,
		Yid: debate.Yid,
		Nid: debate.Nid,
	}

	//将场次信息存储在mysql里面
	if err := dao.Db.Table("debate").Create(&deb).Error; err != nil {
		return http.StatusInternalServerError, errmsg.Error
	}
	now := time.Now().Unix()

	//存储redis
	_, err = dao.Conn.Do(
		"HMSET",
			strconv.Itoa(debateId),
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
