package model


type User struct {
	Uid int32  `gorm:"type:int;not null;primaryKey;auto_increment" json:"uid"`
	//Title string  `gorm:"type:varchar(33);not null" json:"title"`
	Username string   `gorm:"type:varchar(50);DEFAULT:'null'" json:"username" validate:"required,min=6,max=12" label:"用户名"`
	Password string   `gorm:"type:varchar(50);not null" json:"password" validate:"required,min=6,max=18" label:"用户密码"`
	Phone string   `gorm:"type:varchar(30)" json:"phone"`
	Title string  `gorm:"type:varchar(33);not null" json:"title"`
	Email string  `gorm:"type:varchar(33);not null" json:"email" validate:"required,email" label:"邮箱"`
	Img string  `gorm:"type:varchar(33);" json:"img"`
	Role int  `gorm:"type:int;DEFAULT:2;" json:"role"`  //用户的角色，管理员或者非管理员
}


type UserInfo struct {
	Username string   `gorm:"type:varchar(50);not null" json:"username"`
	Score string  `gorm:"type:varchar(33);not null" json:"score"`   //用户的分数存储在redis里面
	Img string  `gorm:"type:varchar(33);" json:"img"`
	Title string  `gorm:"type:varchar(33);not null" json:"title"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email" label:"邮箱"`
	Password string `json:"password" validate:"required,min=6,max=18" label:"用户密码"`
	RememberPassword bool `json:"remember_password"`
}

type UpdateNewPassword struct {
	Email            string `json:"email" validate:"required,email" label:"邮箱"`
	OldPassword      string `json:"old_password" validate:"required,min=6,max=18" label:"用户旧密码"`
	NewPassword      string `json:"new_password" validate:"required,min=6,max=18" label:"用户新密码"`
	CheckNewPassword string `json:"check_new_password"`
}

type RegistryRequest struct {
	Email string  `gorm:"type:varchar(33);not null" json:"email" validate:"required,email" label:"邮箱"`
	Password string   `gorm:"type:varchar(50);not null" json:"password" validate:"required,min=6,max=18" label:"用户密码"`
	Code string `gorm:"type:varchar(10)" json:"code"`
}
