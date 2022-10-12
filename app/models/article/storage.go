package article

import (
	"github.com/lonli7/goblog/pkg/logger"
	"github.com/lonli7/goblog/pkg/model"
	"github.com/lonli7/goblog/pkg/types"
)

func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToUint64(idstr)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

func GetAll() ([]Article, error) {
	var articles []Article
	if err := model.DB.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

func (a *Article) Create() (err error) {
	if err = model.DB.Create(&a).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

func (a *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&a)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

func (a *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&a)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}
