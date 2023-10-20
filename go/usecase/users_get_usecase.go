// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func GetAllUsers() ([]model.User, error) {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	users, err := dao.GetAllUsers()
	if err != nil {
		log.Printf("fail: dao.GetAllUsers(), %v\n", err)
		return nil, err
	}
	return users, nil
}
