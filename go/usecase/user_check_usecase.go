// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func UserCheck(userName string) (model.User, error) {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	user, err := dao.UserCheck(userName)
	if err != nil {
		log.Printf("fail: dao.UserCheck(), %v\n", err)
		return model.User{}, err
	}
	return user, nil
}
