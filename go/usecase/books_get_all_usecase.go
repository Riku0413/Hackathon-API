// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func BooksGetAll() ([]model.Book, error) {
	books, err := dao.BooksGetAll()
	if err != nil {
		log.Printf("fail: dao.BooksGetAll(), %v\n", err)
		return nil, err
	}
	return books, nil
}
