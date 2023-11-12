// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func CommentsGet(itemCategory string, itemId string) ([]model.Comment, error) {
	comments, err := dao.CommentsGet(itemCategory, itemId)
	if err != nil {
		log.Printf("fail: dao.CommentsGet(), %v\n", err)
		return nil, err
	}
	return comments, nil
}
