// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func WorksSearch(key string) ([]model.Work, error) {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	works, err := dao.WorksSearch(key)
	if err != nil {
		log.Printf("fail: dao.WorksSearch(), %v\n", err)
		return nil, err
	}
	return works, nil
}
