package login

import (
	"encoding/base64"
	"github.com/Peterliang233/debate/dao"
	"github.com/Peterliang233/debate/errmsg"
	"github.com/Peterliang233/debate/model"
	"github.com/Peterliang233/debate/router/v1/api/user/login"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
	"net/http"
)

//type Repository interface {
//	Find(id int32) (*model.User, error)
//	Create(*model.User) error
//	Update(*model.User) (*model.User, error)
//}
//
//type User struct {
//	Db *gorm.DB
//}
//
//
////通过id查找该用户是否存在
//func (repo *User) Find(id int32) (*model.User, error) {
//	user := &model.User{}
//	user.Uid = id
//	if err := repo.Db.First(&user).Error; err != nil {
//		return nil, err
//	}
//	return user, nil
//}
//
////更新用户信息
//func (repo *User) Update(user *model.User) (*model.User, error) {
//	if err := repo.Db.Model(&user).Update(user).Error; err != nil {
//		return nil, err
//	}
//	return user, nil
//}
//
////创建用户
//func (repo *User) Create(user *model.User) error {
//	if err := repo.Db.Create(&user).Error; err != nil {
//		return err
//	}
//	return nil
//}


//创建用户
func CreateUser(data *model.User) (int,int) {
	data.Password = ScryptPassword(data.Password)
	if err := dao.Db.Create(data).Error; err != nil {
		return http.StatusInternalServerError,errmsg.Error
	}
	return http.StatusOK,errmsg.Success
}


//注销用户
//func DeleteUser(id int32) int {
//	var user model.User
//	user.Uid = id
//	if err := dao.Db.Where("uid = ",id).Delete(&user).Error; err != nil {
//		return errmsg.Error
//	}
//
//	return errmsg.Success
//}


//修改用户密码
func UpdatePassword(data *login.UpdateNewPassword) (int,int) {
	if data.NewPassword != data.CheckNewPassword {
		return http.StatusBadRequest, errmsg.ErrPasswordDifferent
	}
	var u model.User
	if err := dao.Db.Where("phone = ?", data.Phone).First(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return http.StatusBadRequest, errmsg.ErrPhoneNotExist
		}else{
			return http.StatusBadRequest, errmsg.Error
		}
	}
	if ScryptPassword(data.OldPassword) != u.Password {
		return http.StatusBadRequest, errmsg.ErrPassword
	}
	u.Password = ScryptPassword(data.NewPassword)
	if err := dao.Db.Where("phone = ?", data.Phone).Update("password", u.Password).Error; err != nil {
		return http.StatusInternalServerError, errmsg.Error
	}
	return http.StatusOK, errmsg.Success
}

//手机+密码登录验证
func CheckLogin(login *model.Login) (int,int) {
	var user model.User
	if err := dao.Db.Where("phone = ?", login.Phone).First(&user).Error; err != nil {
		return http.StatusInternalServerError,errmsg.Error
	}
	if ScryptPassword(login.Password) != user.Password {
		return http.StatusBadRequest,errmsg.ErrPassword
	}
	return http.StatusOK,errmsg.Success
}


//修改用户信息
func UpdateUser(u *model.User) int {
	if err := dao.Db.Updates(&u).Where("uid = ?", u.Uid).Error; err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

//密码加密
func ScryptPassword(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{23, 32, 21, 11, 11, 22, 11, 0}
	//加盐
	HashPassword, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(HashPassword)
}


//检查用户名和电话是否存在
func CheckUser(username, phone string) (int,int) {
	var user model.User
	if err := dao.Db.Table("user").Where("username = ?", username).First(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return http.StatusInternalServerError,errmsg.ErrDatabaseFound
		}
		if err = dao.Db.Table("user").Where("phone = ?", phone).First(&user).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return http.StatusInternalServerError,errmsg.ErrDatabaseFound
			}else{
				return http.StatusOK,errmsg.Success
			}
		}else{
			return http.StatusBadRequest,errmsg.ErrUserEmailUsed
		}
	}else{
		return http.StatusBadRequest,errmsg.ErrUserNameUsed
	}
}