package login

import (
	"github.com/Peterliang233/debate/dao"
	"github.com/Peterliang233/debate/errmsg"
	"github.com/garyburd/redigo/redis"
	"log"
	"math/rand"
	"strconv"
	"time"
)

//发送邮件
func SendEmail(Email string) (string, int) {
	code := GetCode()
	//e := email.NewEmail()
	//e.From = "Peterliang <ncuyanping666@126.com>"
	//e.To = []string{Email}
	//e.Subject = "注册通知"
	//e.Text = []byte("              注册验证通知\n您好！\n您的邮箱在" +
	//	time.Now().Format("2006-01-02 15:04:05") +
	//	"被用于注册\"来辩\"\n验证码为："+code +"\n五分钟内有效.")
	//err := e.Send(config.EmailSetting.Addr, smtp.PlainAuth("", config.EmailSetting.Username,
	//	config.EmailSetting.Password, config.EmailSetting.Host))
	//if err != nil {
	//	log.Fatal(err)
	//	return "", errmsg.Error
	//}
	return code, errmsg.Success
}


//将验证码在redis里面存储5分钟
func SetRedis(email, code string) bool {
	//conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil {
	//	log.Println("connect redis error :", err)
	//	return false
	//}
	//defer conn.Close()
	_, err := dao.Conn.Do("SET", email, code)  //把code放到email里面
	if err != nil {
		log.Println("redis set error:", err)
		return false
	}
	_, err = dao.Conn.Do("expire", email, 300)  //放到redis里面缓存5分钟
	if err != nil {
		log.Println("set expire error: ", err)
		return false
	}
	return true
}

func GetRedis(email string) string {
	//conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil {
	//	log.Println("connect redis error :", err)
	//}
	//defer conn.Close()
	code, err := redis.String(dao.Conn.Do("GET", email))
	if err != nil {
		log.Println("redis get error:", err)
	}
	return code
}


//获取生成验证码
func GetCode() string {
	rand.Seed(time.Now().UnixNano())

	code := rand.Intn(899999) + 100000

	return strconv.Itoa(code)
}


//验证验证码
func Validation(email, code string) bool {
	if code == GetRedis(email) {
		return true
	}
	return false
}