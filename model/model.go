package model


type User struct {
	Uid int32  `gorm:"type:int;not null;primaryKey;auto_increment" json:"uid"`
	//Title string  `gorm:"type:varchar(33);not null" json:"title"`
	Username string   `gorm:"type:varchar(50);not null" json:"username" validate:"required,min:6,max=12" label:"用户名"`
	Password string   `gorm:"type:varchar(50);not null" json:"password" validate:"required,min:6,max=18" label:"用户密码"`
	Phone string   `gorm:"type:varchar(30);not null" json:"phone" validate:"required,phone" label:"电话"`
	Email string  `gorm:"type:varchar(33)" json:"email" validate:"required,email" label:"邮箱"`
	Role int  `gorm:"type:int;DEFAULT:2;" json:"role"`  //用户的角色，管理员或者非管理员
	Rid int  `gorm:"type:int;DEFAULT:1" json:"rid"`   //关联用户头衔的id
	Score int `gorm:"type:int;DEFAULT:0" json:"score"`
}

type Identify struct {
	Id int32  `gorm:"type:int;not null;primaryKey;auto_increment" json:"id"`
	Name string   `gorm:"type:varchar(33)" json:"name"`
}

type Login struct {
	Phone string  `json:"phone"`
	Password string 	`json:"password"`
}



