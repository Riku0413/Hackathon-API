// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func VideoGet(videoId string) (model.Video, error) {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	video, err := dao.VideoGet(videoId)
	if err != nil {
		log.Printf("fail: dao.VideoGet(), %v\n", err)
		return model.Video{}, err
	}
	return video, nil
}
