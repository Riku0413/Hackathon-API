// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func LikeCheck(category string, itemId string, userId string) (model.Like, error) {
	like, err := dao.LikeCheck(category, itemId, userId)
	if err != nil {
		log.Printf("fail: dao.LikeCheck(), %v\n", err)
		return model.Like{}, err
	}
	return like, nil
}
