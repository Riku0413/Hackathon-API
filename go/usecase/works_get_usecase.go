// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func WorksGet(userId string) ([]model.Work, error) {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	works, err := dao.WorksGet(userId)
	if err != nil {
		log.Printf("fail: dao.WorksGet(), %v\n", err)
		return nil, err
	}
	return works, nil
}
