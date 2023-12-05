// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func WorksGetAll() ([]model.Work, error) {
	works, err := dao.WorksGetAll()
	if err != nil {
		log.Printf("fail: dao.WorksGetAll(), %v\n", err)
		return nil, err
	}
	return works, nil
}
