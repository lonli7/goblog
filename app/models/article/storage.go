package article

import (
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