# School API Gateway

動かし方メモ

```
$ asdf plugin list all | grep golang
golang                       *https://github.com/asdf-community/asdf-golang.git
golangci-lint                *https://github.com/hypnoglow/asdf-golangci-lint.git
$ asdf plugin add golang https://github.com/asdf-community/asdf-golang.git
$ asdf install golang latest
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
$ asdf install golang 1.20
$ asdf global golang 1.20

$ go run main.go
2024/06/01 12:10:46 gRPC を起動 0.0.0.0:8080
2024/06/01 12:10:46 gRPC-Gateway 起動 http://0.0.0.0:8090

$ curl -X POST http://localhost:8090/v1/school/labos
{"labos":[{"name":"研究室-0001"},{"name":"研究室-0002"},{"name":"研究室-0003"},{"name":"研究室-0004"},{"name":"研究室-0005"}],"offset":5}
```
