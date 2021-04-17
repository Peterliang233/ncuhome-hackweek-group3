package debate

import (
	"fmt"
	"github.com/Peterliang233/debate/dao"
	"github.com/Peterliang233/debate/errmsg"
	"github.com/Peterliang233/debate/model"
	"github.com/garyburd/redigo/redis"
	"log"
	"net/http"
	"strconv"
)


//将辩论内容记录在redis
func CreateRecord(debate * model.DebateRedis) (int,int) {
	debateId, err := redis.Int(dao.Conn.Do("incr", "debateId"))  //记录下辩论的总场次
	if err != nil {
		log.Fatalf(err.Error())
		return http.StatusInternalServerError, errmsg.Error
	}

	deb := model.DebateContent{
		Id: debateId,
		Title: debate.Title,
		PositiveUsername: debate.PositiveContent,
		NegativeUsername: debate.NegativeContent,
		BeginTime: debate.BeginTime,
	}

	//将场次信息存储在mysql里面
	if err := dao.Db.Table("debate").Create(&deb).Error; err != nil {
		fmt.Print(err)
		return http.StatusInternalServerError, errmsg.Error
	}
	//存储redis
	_, err = dao.Conn.Do(
		"HMSET",
			strconv.Itoa(debateId),
			"title", debate.Title,
			"positive_content", debate.PositiveContent,
			"negative_content", debate.NegativeContent,
			"positive_username", debate.PositiveUsername,
			"negative_username", debate.NegativeUsername,
			"time", debate.BeginTime,
	)
	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, errmsg.Error
	}
	return http.StatusOK, errmsg.Success
}

//通过id获取某一个场次的辩论记录
func GetRedisHashRecord(id string) (map[string]string, int, int) {
	resKey, err := redis.Values(dao.Conn.Do("hkeys", id))
	if err != nil {
		fmt.Println("hkeys failed", err.Error())
		return nil, http.StatusInternalServerError, errmsg.Error
	}
	resValue, err := redis.Values(dao.Conn.Do("hvals", id))
	if err != nil {
		fmt.Println("hvals failed", err.Error())
		return nil, http.StatusInternalServerError, errmsg.Error
	}
	var s1,s2 []string
	for _, v := range resKey {
		s1 = append(s1, string(v.([]byte)))
		//fmt.Printf("%s", v.([]byte))
	}
	for _, v := range resValue {
		s2 = append(s2, string(v.([]byte)))
		//fmt.Printf("%s", v.([]byte))
	}


	result := make(map[string]string)
	for i := 0; i < len(s1); i ++ {
		result[s1[i]] = s2[i]
	}
	return result, http.StatusOK, errmsg.Success
}


//更新为选择正方
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


//更新为选择反方
func UpdateNegative(content * model.DebateContent) (StatusCode, code int){
	if err := dao.Db.Table("debate").Where("title = ?", content.Title).
		Update("negative_username", content.NegativeUsername).Error; err != nil {
		return http.StatusInternalServerError, errmsg.Error
	}
	return http.StatusOK, errmsg.Success
}


//分页展示
func GetRecords(page model.Page) (records []model.DebateContent, statusCode, code int) {
	if err := dao.Db.Table("debate").
		Limit(page.PageSize).Offset((page.PageNum-1)*page.PageSize).
		Find(&records).Error; err != nil {
		return nil, http.StatusInternalServerError, errmsg.Error
	}
	return records, http.StatusOK, errmsg.Success
}

//func GetRecordsByTime(page model.Page) (records []model.DebateContent, statusCode, code int) {
//	if err := dao.Db.Table("debate").Order("CreateAt").Error; err != nil {
//
//	}
//}