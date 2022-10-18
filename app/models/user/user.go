package user

import (
	"github.com/lonli7/goblog/app/models"
	"github.com/lonli7/goblog/pkg/password"
)

type User struct {
	models.BaseModel

	Name     string `gorm:"column:name;type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"column:email;type:varchar(255);default:NULL;unique" valid:"email"`
	Password string `gorm:"column:password;type:varchar(255)not null" valid:"password"`
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

func (user *User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, user.Password)
}
