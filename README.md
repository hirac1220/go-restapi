# Go-REST APIテストプロ

## Docker
### 環境構築(Docker)
事前準備:
* nuxt.config.jsのURL/Portをローカルへ変更
* DB接続先をDockerのMySQLへ変更

開発環境を起動
```
docker-compose up
```

apiサーバをビルド
```
docker exec -it {docker_id} sh

# /go/src/appで
go build
```
apiサーバを起動
```
./api
```
web開発環境へ移動し起動
```
cd {path}/go-restapi/web
npm run dev
```
### 動作確認(Docker)

ローカル環境でのWebサイト  
[Memo](http://localhost:3000/memo)

参照
```
curl http://localhost:18080/memo
```
追加
```
curl -X POST -d '{"text":"test"}' http://localhost:18080/memo
```
更新
```
curl -X PUT -d '{"text":"testtest"}' http://localhost:18080/memo/1
```
削除
```
curl -X DELETE http://localhost:18080/memo/1
```

## Google App Engine
### 環境構築(GAE)

事前準備:
* nuxt.config.jsのURL/PortをGAEへ変更
* DB接続先をCloud SQLへ変更

API: 開発環境へ移動しビルド
```
cd {path}/go-restapi/api
go build
```
Web: 開発環境へ移動しビルド
```
cd {path}/go-restapi/web
npm run build
```
GCPへログインし、認証コード貼り付ける(初回のみ)
```
gcloud auth login
```
GAEへデプロイ(web/api各ディレクトリより)
```
gcloud app deploy --project {PROJECT_ID}
```
### 動作確認(GAE)

Webサイト on GAE  
[Memo](https://{WEB_SERVICE}-dot-{PROJECT_ID}.appspot.com/)

参照
```
curl https://{API_SERVICE}-dot-{PROJECT_ID}.appspot.com/memo
```
追加
```
curl -X POST -d '{"text":"test"}' https://{API_SERVICE}-dot-{PROJECT_ID}.appspot.com/memo
```
更新
```
curl -X PUT -d '{"text":"testtest"}' https://{API_SERVICE}-dot-{PROJECT_ID}.appspot.com/memo/1
```
削除
```
curl -X DELETE https://{SERVICE}-dot-{API_PROJECT_ID}.appspot.com/memo/1
```
