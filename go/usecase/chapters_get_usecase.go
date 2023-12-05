// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func ChaptersGet(bookId string) ([]model.Chapter, error) {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	books, err := dao.ChaptersGet(bookId)
	if err != nil {
		log.Printf("fail: dao.ChaptersGet(), %v\n", err)
		return nil, err
	}
	return books, nil
}
