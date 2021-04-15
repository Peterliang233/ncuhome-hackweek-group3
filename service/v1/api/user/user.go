package user

import (
	"github.com/Peterliang233/debate/dao"
	"github.com/Peterliang233/debate/errmsg"
	"github.com/Peterliang233/debate/model"
	"github.com/Peterliang233/debate/service/v1/api/user/login"
	"github.com/jinzhu/gorm"
	"net/http"
)

//根据查询用户信息
func GetUserInfo(email string) (int,int, *model.UserInfo) {
	//fmt.Printf(email)
	var u model.User
	if err := dao.Db.Table("user").Where("email = ?", email).First(&u).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return http.StatusInternalServerError,errmsg.Error,nil
		}
		return http.StatusNotFound, errmsg.Error,nil
	}
	userInfo := model.UserInfo{
		Username: u.Username,
		Img: u.Img,
	}
	//userInfo.Score = dao.Conn.Do("GET", string(u.Uid))
	return http.StatusOK, errmsg.Success, &userInfo
}

//修改用户密码
func UpdatePassword(data *model.UpdateNewPassword) (int,int) {
	if data.NewPassword != data.CheckNewPassword {
		return http.StatusBadRequest, errmsg.ErrPasswordDifferent
	}
	var u model.User
	if err := dao.Db.Table("user").Where("email = ?", data.Email).First(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return http.StatusBadRequest, errmsg.ErrEmailNotExist
		}else{
			return http.StatusBadRequest, errmsg.Error
		}
	}
	if login.ScryptPassword(data.OldPassword) != u.Password {
		return http.StatusBadRequest, errmsg.ErrPassword
	}
	u.Password = login.ScryptPassword(data.NewPassword)
	if err := dao.Db.Table("user").Where("email = ?", data.Email).Update("password", u.Password).Error; err != nil {
		return http.StatusInternalServerError, errmsg.Error
	}
	return http.StatusOK, errmsg.Success
}


//修改用户信息(username和手机号码)
func UpdateUserInfo(u *model.User) (int,int) {
	//检查用户名是否被使用
	StatusCode, code := login.CheckUsername(u.Username)
	if code != errmsg.Success {
		return StatusCode, code
	}
	if err := dao.Db.Table("user").Where("email = ?", u.Email).Updates(map[string]interface{}{
		"username": u.Username,"phone": u.Phone,
	}).Error; err != nil {
		return http.StatusInternalServerError, errmsg.Error
	}
	return http.StatusOK, errmsg.Success
}