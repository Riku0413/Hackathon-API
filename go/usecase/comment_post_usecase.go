// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func CommentPost(comment model.Comment) error {
	if err := dao.CommentPost(comment); err != nil {
		log.Printf("fail: dao.CommentPost, %v\n", err)
		return err
	}
	return nil
}
