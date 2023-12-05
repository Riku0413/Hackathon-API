// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func VideosGetAll() ([]model.Video, error) {
	videos, err := dao.VideosGetAll()
	if err != nil {
		log.Printf("fail: dao.VideosGetAll(), %v\n", err)
		return nil, err
	}
	return videos, nil
}
