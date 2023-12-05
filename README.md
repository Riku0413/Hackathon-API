# Hackathon-API

## 概要

自作したナレッジベース（Webアプリ）のバックエンド

## 実行方法（MySQLコンテナの作成）

1. リポジトリのクローン
   ```
   git clone "https://github.com/Riku0413/Hackathon-API.git"
   ```
2. ディレクトリ移動
   ```
   cd ./mysql
   ```
3. アクセス権の追加
   ```
   chmod a+x ./init/*.sh
   ```

4. コンテナの起動
   ```
   docker-compose up -d
   ```

## 実行方法（Goサーバーの立ち上げ）

1. リポジトリのクローン
   ```
   git pull "https://github.com/Riku0413/Hackathon-API.git"
   ```
2. ディレクトリ移動
   ```
   cd ./go
   ```
3. サーバーの起動
   ```
   go run ./...
   ```
   
## 搭載機能

- 記事、本、動画、プロダクトという４種類のアイテム
- アイテムの一覧表示
- アイテムの３通りのソート
- アイテムの追加、編集、削除
- アイテムへのコメント投稿
- フリーワード検索
- サインアップ、サインイン、サインアウト
- プロフィール編集

## ディレクトリ構成
```
go
├── controller
├── usecase
├── dao
└── model
```

## 技術スタック

- Go
- SQL
- Cloud Run（過去のデプロイ時に使用）
- Cloud SQL（過去のデプロイ時に使用）
