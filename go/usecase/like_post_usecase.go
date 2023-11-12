// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func LikePost(like model.Like) error {
	if err := dao.LikePost(like); err != nil {
		log.Printf("fail: dao.LikePost, %v\n", err)
		return err
	}
	return nil
}
