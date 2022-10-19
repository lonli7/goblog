package category

import (
	"github.com/lonli7/goblog/pkg/logger"
	"github.com/lonli7/goblog/pkg/model"
	"github.com/lonli7/goblog/pkg/types"
)

func (c *Category) Create() (err error){
	if err = model.DB.Create(&c).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

func All() ([]Category, error) {
	var categories []Category
	if err := model.DB.Order("created_at desc").Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

func Get(cid string) (Category, error) {
	var category Category
	id := types.StringToUint64(cid)
	if err := model.DB.First(&category, id).Error; err != nil {
		return category, err
	}
	return category, nil
}
