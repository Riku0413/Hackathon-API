// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"log"
)

func WorkDelete(workId string) error {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	err := dao.WorkDelete(workId)
	if err != nil {
		log.Printf("fail: dao.WorkDelete(), %v\n", err)
		return err
	}
	return nil
}
