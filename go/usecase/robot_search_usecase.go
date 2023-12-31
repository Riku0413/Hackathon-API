// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func GetRobotsByName(name string) ([]model.Robot, error) {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	// ここは、本当は1行で書けるけど、可読性確保のためにわざと5行で書いている
	// usecaseでは、Go特有の型、以外は登場させない！！
	robots, err := dao.GetRobotsByName(name)
	if err != nil {
		log.Printf("fail: dao.GetRobotsByName, %v\n", err)
		return nil, err
	}
	return robots, nil
}
