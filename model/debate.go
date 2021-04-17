package model


//存储在redis
type DebateRedis struct {
	Id int32 `gorm:"type:int;not null;auto_increment" json:"id"`
	Yid int32 `gorm:"type:int;not null" json:"yid" label:"正方id"`
	Nid int32 `gorm:"type:int;not null" json:"nid" label:"反方id"`
	Title string `gorm:"type:varchar(33);not null" json:"title" label:"辩论标题"`
	NegativeContent string  `json:"negative_content" label:"正方发言"`
	PositiveContent string `json:"positive_content" label:"反方发言"`
}


//存储在mysql，便于后面查询单个用户的辩论情况
type DebateMysql struct {
	Id int32 `gorm:"type:int;not null;auto_increment" json:"id"`
	Yid int32 `gorm:"type:int;not null" json:"yid" label:"正方id"`
	Nid int32 `gorm:"type:int;not null" json:"nid" label:"反方id"`
}


//辩论的情况
type DebateContent struct {
	Title string  `gorm:"type:varchar(100);not null" json:"title" label:"辩题"`
	Positive string `gorm:"type:varchar(20);"`
	Negative string
}


//进行辩论请求
type DebateRequest struct {
	RoomID string `json:"room_id"`
	UserID string `json:"user_id"`
	UseName string `json:"use_name"`
}