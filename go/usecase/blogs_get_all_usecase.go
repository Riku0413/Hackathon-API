// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func BlogsGetAll() ([]model.Blog, error) {
	blogs, err := dao.BlogsGetAll()
	if err != nil {
		log.Printf("fail: dao.BlogsGetAll(), %v\n", err)
		return nil, err
	}
	return blogs, nil
}
