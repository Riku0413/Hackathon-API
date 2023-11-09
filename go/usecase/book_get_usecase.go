// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func BookGet(bookId string) (model.Book, error) {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	book, err := dao.BookGet(bookId)
	if err != nil {
		log.Printf("fail: dao.BookGet(), %v\n", err)
		return model.Book{}, err
	}
	return book, nil
}
