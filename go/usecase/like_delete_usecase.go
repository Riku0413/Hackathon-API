// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"log"
)

func LikeDelete(likeId string, itemCategory string, itemId string) error {
	err := dao.LikeDelete(likeId, itemCategory, itemId)
	if err != nil {
		log.Printf("fail: dao.LikeDelete(), %v\n", err)
		return err
	}
	return nil
}
