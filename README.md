# School API Gateway

## ローカルの開発環境の設定

```bash
# asdf のインストール
$ brew install coreutils curl git
$ brew install asdf
$ echo -e "\n. $(brew --prefix asdf)/libexec/asdf.sh" >> ${ZDOTDIR:-~}/.zshrc

# golang プラグインのインストール
$ asdf plugin list all | grep golang
golang                       *https://github.com/asdf-community/asdf-golang.git
golangci-lint                *https://github.com/hypnoglow/asdf-golangci-lint.git
$ asdf plugin add golang https://github.com/asdf-community/asdf-golang.git

# golang のインストール
$ asdf install golang 1.22.4
Platform 'darwin' supported!
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 64.1M  100 64.1M    0     0  28.0M      0  0:00:02  0:00:02 --:--:-- 28.0M
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    64  100    64    0     0    356      0 --:--:-- --:--:-- --:--:--   357
verifying checksum
/Users/yamagame/.asdf/downloads/golang/1.22.3/archive.tar.gz: OK
checksum verified
$ asdf local golang 1.22.4
```

## サーバーの起動

```bash
$ go run main.go
2024/06/01 12:10:46 gRPC を起動 0.0.0.0:8080
2024/06/01 12:10:46 gRPC-Gateway 起動 http://0.0.0.0:8090

$ curl -X POST http://localhost:8090/v1/school/labos
{"labos":[{"name":"研究室-0001"},{"name":"研究室-0002"},{"name":"研究室-0003"},{"name":"研究室-0004"},{"name":"研究室-0005"}],"offset":5}
```

## Docker の操作

```sh
# コンテナの起動
$ docker-compose up -d

# コンテナの削除
$ docker-compose down --rmi 

# イメージも含めての削除
$ docker-compose down --rmi all

# go-dev コンテナにログイン
$ docker exec -it school-api-gateway-go-dev-1 bash
```

## SQL Tools プラグインの使い方

[SQLTools](https://marketplace.visualstudio.com/items?itemName=mtxr.sqltools) をインストール

以下、setting.json に記入

```json
{
            :
  "sqltools.format": {
    "reservedWordCase": "upper",
    "linesBetweenQueries": "preserve"
  }
}
```

下記設定で school-database に接続

- Connection name: school-database
- Server Address: localhost
- Port: 3336
- Database: school-database
- Usrname: root

SQLを選択して、⌘E ⌘Eで選択実行。


## Grom で接続

- [Go言語のための最高のORMライブラリ](https://gorm.io/ja_JP/)

```golang
func sample() {
  // データベースに接続
	db := infra.DB()

  // Labosテーブルの作成
	db.AutoMigrate(&repository.Labos{})

  // 研究室-001 の作成
	db.Create(&repository.Labos{
		Name: "研修室-001",
	})

  // Labosテーブルを検索
  var labos []*repository.Labos
  db.Limit(10).Offset(0).Find(&labos) 
  // 検索結果をJSON文字列で表示
	jsonData, _ := json.MarshalIndent(labos, "", "  ")
	fmt.Printf("%s\n", jsonData)
}
```
