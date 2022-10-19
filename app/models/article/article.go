package article

import (
	"github.com/lonli7/goblog/app/models"
	"github.com/lonli7/goblog/pkg/route"
	"strconv"
)

type Article struct {
	models.BaseModel

	Title string `gorm:"column:title;type:varchar(255);not null" valid:"title"`
	Body string `gorm:"column:body;type:varchar(255);not null" valid:"body"`
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(a.ID, 10))
}
