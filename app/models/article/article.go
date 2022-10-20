package article

import (
	"goblog/app/models"
	"goblog/app/models/category"
	"goblog/app/models/user"
	"goblog/pkg/route"
	"strconv"
)

type Article struct {
	models.BaseModel

	Title      string `gorm:"column:title;type:varchar(255);not null" valid:"title"`
	Body       string `gorm:"column:body;type:varchar(255);not null" valid:"body"`
	UserID     uint64 `gorm:"not null:index"`
	User       user.User
	CategoryId uint64 `gorm:"not null:index"`
	Category   category.Category
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(a.ID, 10))
}

func (a Article) CreatedAtDate() string {
	return a.CreatedAt.Format("2006-01-02")
}
