// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func WorkGet(workId string) (model.Work, error) {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	work, err := dao.WorkGet(workId)
	if err != nil {
		log.Printf("fail: dao.WorkGet(), %v\n", err)
		return model.Work{}, err
	}
	return work, nil
}
