// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func BlogsSearch(key string) ([]model.Blog, error) {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	blogs, err := dao.BlogsSearch(key)
	if err != nil {
		log.Printf("fail: dao.BlogsSearch(), %v\n", err)
		return nil, err
	}
	return blogs, nil
}
