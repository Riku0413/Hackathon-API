// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func BlogGet(blogId string) (model.Blog, error) {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	blog, err := dao.BlogGet(blogId)
	if err != nil {
		log.Printf("fail: dao.BlogGet(), %v\n", err)
		return model.Blog{}, err
	}
	return blog, nil
}
