// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func ChapterGet(chapterId string) (model.Chapter, error) {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	chapter, err := dao.ChapterGet(chapterId)
	if err != nil {
		log.Printf("fail: dao.ChapterGet(), %v\n", err)
		return model.Chapter{}, err
	}
	return chapter, nil
}
