package category

import (
	"github.com/lonli7/goblog/app/models"
	"github.com/lonli7/goblog/pkg/route"
)

type Category struct {
	models.BaseModel

	Name string `gorm:"column:name;type:varchar(255);not null;" valid:"name"`
}

func (c Category) Link() string {
	return route.Name2URL("categories.show", "id", c.GetStringID())
}
