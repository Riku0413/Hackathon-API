// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func RegisterRobot(robot model.Robot) error {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	// 本当は、ここのreturn文は1行にまとめられるけど、可読性向上のために2行に分けておく
	if err := dao.RegisterRobot(robot); err != nil {
		log.Printf("fail: dao.RegisterRobot, %v\n", err)
		return err
	}
	return nil
}
