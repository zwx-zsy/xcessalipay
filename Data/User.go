package Data

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserId      string    `json:"user_id"`
	Email       string    `json:"email"`
	UserName    string    `json:"userName"`
	MobilePhone string    `json:"mobilePhone"`
	PassWord    string    `json:"passWord"`
	LogInTime   time.Time `json:"logInTime"`
}

func (u *User) Insert() (result *gorm.DB) {
	result = _db.Model(&User{}).Create(&u)
	return
}
