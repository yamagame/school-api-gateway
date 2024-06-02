package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pbSchool "github.com/yamagame/school-api-gateway/proto/school"
)

// School サービスの構造体
type server struct{}

// ListLabos 研究室の一覧を返す
func (*server) ListLabos(_ context.Context, in *pbSchool.ListLabosRequest) (*pbSchool.ListLabosResponse, error) {
	pageSize := int32(5)
	if in.PageSize != nil {
		pageSize = *in.PageSize
	}
	offset := int32(0)
	if in.Offset != nil {
		offset = *in.Offset
	}
	labos := []*pbSchool.Labo{}
	for i := int32(0); i < pageSize; i++ {
		labos = append(labos, &pbSchool.Labo{
			Name: fmt.Sprintf("研究室-%04d", i+1+offset),
		})
	}
	return &pbSchool.ListLabosResponse{Labos: labos, Offset: pageSize + offset}, nil
}

func main() {
	// TCP ポートを作成する
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("TCP ポートのリッスンに失敗:", err)
	}

	// gRPC サーバーオブジェクトを作成する
	s := grpc.NewServer()
	// School サービスを接続
	pbSchool.RegisterShoolServer(s, &server{})
	// gRPC サーバーを起動
	log.Println("gRPC を起動 locahost:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// gRPC サーバーに接続するクライアントの作成
	conn, err := grpc.NewClient(
		"localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("gRPC クライアントの起動に失敗:", err)
	}

	// ServeMux の作成
	gwmux := runtime.NewServeMux()
	// School サービスを gRPC-Gateway に登録
	err = pbSchool.RegisterShoolHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("サービスの登録に失敗:", err)
	}

	// HTTP サーバーを作成
	gwServer := &http.Server{
		Addr:    "localhost:8090",
		Handler: gwmux,
	}

	log.Println("gRPC-Gateway 起動 http://locahost:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
