package login

import (
	"github.com/Peterliang233/debate/config"
	"github.com/Peterliang233/debate/errmsg"
	"github.com/garyburd/redigo/redis"
	"github.com/jordan-wright/email"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

//发送邮件
func SendEmail(Email string) (string, int) {
	code := GetCode()
	e := email.NewEmail()
	e.From = "Peterliang <ncuyanping666@126.com>"
	e.To = []string{Email}
	e.Subject = "注册通知"
	e.Text = []byte("注册验证通知\n您好！\n您的邮箱在" +
		time.Now().Format("2006-01-02 15:04:05") + "被用于注册.\n验证码为："+code +"\n五分钟内有效.")
	err := e.Send(config.EmailSetting.Addr, smtp.PlainAuth("", config.EmailSetting.Username,
		config.EmailSetting.Password, config.EmailSetting.Host))
	if err != nil {
		return "", errmsg.Error
	}
	return code, errmsg.Success
}

func SetRedis(code string) bool {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println("connect redis error :", err)
		return false
	}
	defer conn.Close()
	_, err = conn.Do("SET", "email", code)
	if err != nil {
		log.Println("redis set error:", err)
		return false
	}
	_, err = conn.Do("expire", "email", 300)
	if err != nil {
		log.Println("set expire error: ", err)
		return false
	}
	return true
}

func GetRedis(userId string) string {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println("connect redis error :", err)
	}
	defer conn.Close()
	code, err := redis.String(conn.Do("GET", userId))
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
func Validation(code string) bool {
	if code == GetRedis("email") {
		return true
	}
	return false
}