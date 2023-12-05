// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"log"
)

func BlogDelete(blogId string) error {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	err := dao.BlogDelete(blogId)
	if err != nil {
		log.Printf("fail: dao.BlogDelete(), %v\n", err)
		return err
	}
	return nil
}
