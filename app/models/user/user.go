package user

import "github.com/lonli7/goblog/app/models"

type User struct {
	models.BaseModel

	Name     string `gorm:"column:name;type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"column:email;type:varchar(255);default:NULL;unique" valid:"email"`
	Password string `gorm:"column:password;type:varchar(255)not null" valid:"password"`
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}
