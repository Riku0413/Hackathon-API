package model

import (
	"github.com/oklog/ulid/v2"
	_ "github.com/oklog/ulid/v2"
	"math/rand"
	"time"
)

func generateTimestamp() (string, error) {
	now := time.Now()
	seconds := now.Unix() // 秒単位のタイムスタンプ
	t := time.Unix(seconds, 0)
	return t.Format("2006-01-02 15:04:05"), nil // フォーマットを指定
}

func generateULID() (ulid.ULID, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	ulidValue, err := ulid.New(ms, entropy)
	return ulidValue, err
}
