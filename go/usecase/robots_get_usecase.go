// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func GetAllRobots() ([]model.Robot, error) {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	robots, err := dao.GetAllRobots()
	if err != nil {
		log.Printf("fail: dao.GetAllRobots(), %v\n", err)
		return nil, err
	}
	return robots, nil
}
