// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func BlogsGet(userId string) ([]model.Blog, error) {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	blogs, err := dao.BlogsGet(userId)
	if err != nil {
		log.Printf("fail: dao.BlogsGet(), %v\n", err)
		return nil, err
	}
	return blogs, nil
}
