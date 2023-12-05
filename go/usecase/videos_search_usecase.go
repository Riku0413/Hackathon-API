// usecaseは入力データが正当であることを前提として、ビジネスロジックを実行する
package usecase

import (
	"github.com/Riku0413/Hackathon-API/dao"
	"github.com/Riku0413/Hackathon-API/model"
	"log"
)

func VideosSearch(key string) ([]model.Video, error) {
	// ここでは何もせずに、daoにバトンを渡す → 返ってきたデータをcontrollerに返す
	videos, err := dao.VideosSearch(key)
	if err != nil {
		log.Printf("fail: dao.VideosSearch(), %v\n", err)
		return nil, err
	}
	return videos, nil
}
