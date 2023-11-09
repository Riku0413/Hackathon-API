// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"log"
)

func VideoDelete(videoId string) error {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	err := dao.VideoDelete(videoId)
	if err != nil {
		log.Printf("fail: dao.VideoDelete(), %v\n", err)
		return err
	}
	return nil
}
