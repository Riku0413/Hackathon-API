// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func BooksGet(userId string) ([]model.Book, error) {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	books, err := dao.BooksGet(userId)
	if err != nil {
		log.Printf("fail: dao.BooksGet(), %v\n", err)
		return nil, err
	}
	return books, nil
}
