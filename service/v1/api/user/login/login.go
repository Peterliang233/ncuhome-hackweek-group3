package login

import (
	"encoding/base64"
	"github.com/Peterliang233/debate/dao"
	"github.com/Peterliang233/debate/errmsg"
	"github.com/Peterliang233/debate/model"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
	"net/http"
)

//type Repository interface {
//	Find(id int32) (*dao.User, error)
//	Create(*dao.User) error
//	Update(*dao.User) (*dao.User, error)
//}
//
//type User struct {
//	Db *gorm.DB
//}
//
//
////通过id查找该用户是否存在
//func (repo *User) Find(id int32) (*dao.User, error) {
//	user := &dao.User{}
//	user.Uid = id
//	if err := repo.Db.First(&user).Error; err != nil {
//		return nil, err
//	}
//	return user, nil
//}
//
////更新用户信息
//func (repo *User) Update(user *dao.User) (*dao.User, error) {
//	if err := repo.Db.Model(&user).Update(user).Error; err != nil {
//		return nil, err
//	}
//	return user, nil
//}
//
////创建用户
//func (repo *User) Create(user *dao.User) error {
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
//	var user dao.User
//	user.Uid = id
//	if err := dao.Db.Where("uid = ",id).Delete(&user).Error; err != nil {
//		return errmsg.Error
//	}
//
//	return errmsg.Success
//}

//邮箱+密码登录验证
func CheckLogin(login *model.Login) (int,int) {
	var user model.User
	if err := dao.Db.Where("email = ?", login.Email).First(&user).Error; err != nil {
		return http.StatusInternalServerError,errmsg.Error
	}
	if ScryptPassword(login.Password) != user.Password {
		return http.StatusBadRequest,errmsg.ErrPassword
	}
	return http.StatusOK,errmsg.Success
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


//检查邮箱是否存在
func CheckEmail(email string) (int,int) {
	var user model.User
	if err := dao.Db.Table("user").Where("email = ?", email).First(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return http.StatusInternalServerError,errmsg.ErrDatabaseFound
		}else{
			return http.StatusOK,errmsg.Success
		}
	}else{
		return http.StatusBadRequest,errmsg.ErrUserEmailUsed
	}
}


//检查用户名是否存在
func CheckUsername(username string) (int,int) {
	var user model.User
	//检查邮箱是否存在
	if err := dao.Db.Table("user").Where("username = ?", username).First(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return http.StatusInternalServerError,errmsg.ErrDatabaseFound
		}else{
			return http.StatusOK,errmsg.Success
		}
	}else{
		return http.StatusBadRequest,errmsg.ErrUserNameUsed
	}
}

