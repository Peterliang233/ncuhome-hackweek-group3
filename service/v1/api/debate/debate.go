package debate

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
			"yid", strconv.Itoa(int(debate.Yid)),
			"nid", strconv.Itoa(int(debate.Nid)),
			"time", now,
	)
	if err != nil {
		return http.StatusInternalServerError, errmsg.Error
	}
	return http.StatusOK, errmsg.Success
}

//通过id获取某一个场次的辩论记录
func GetRedisHashRecord(id string) (interface{}, int, int) {
	result, err := redis.Values(dao.Conn.Do("HGETALL", id))
	if err != nil {
		return nil, http.StatusInternalServerError, errmsg.Error
	}
	return result, http.StatusOK, errmsg.Success
}



func UpdatePositive(content * model.DebateContent) (StatusCode, code int){
	//var user model.User
	////获取用户id
	//if err := dao.Db.Table("user").Where("username = ?", content.NegativeUsername).
	//	First(&user).Error; err != nil {
	//	return http.StatusInternalServerError, errmsg.Error
	//}
	if err := dao.Db.Table("debate").Where("title = ?", content.Title).
		Update("positive_username", content.PositiveUsername).Error; err != nil {
		return http.StatusInternalServerError, errmsg.Error
	}
	return http.StatusOK, errmsg.Success
}

func UpdateNegative(content * model.DebateContent) (StatusCode, code int){
	//var user model.User
	////获取用户id
	//if err := dao.Db.Table("user").Where("username = ?", content.NegativeUsername).
	//	First(&user).Error; err != nil {
	//	return http.StatusInternalServerError, errmsg.Error
	//}
	if err := dao.Db.Table("debate").Where("title = ?", content.Title).
		Update("negative_username", content.NegativeUsername).Error; err != nil {
		return http.StatusInternalServerError, errmsg.Error
	}
	return http.StatusOK, errmsg.Success
}

func GetRecords(records []model.DebateContent) (statusCode, code int) {
	if err := dao.Db.Find(&records).Error; err != nil {
		return http.StatusInternalServerError, errmsg.Error
	}
	return http.StatusOK, errmsg.Success
}