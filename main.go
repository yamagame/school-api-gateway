package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/yamagame/school-api-gateway/infra"
	"github.com/yamagame/school-api-gateway/infra/repository"
	pbSchool "github.com/yamagame/school-api-gateway/proto/school"
	"github.com/yamagame/school-api-gateway/service/svcconv"
)

var ServerHost = "localhost"

func init() {
	serverhost, ok := os.LookupEnv("SERVER_HOST")
	if ok {
		ServerHost = serverhost
	}
}

// School サービスの構造体
type server struct {
	Repo *repository.School
}

// ListLabos 研究室の一覧を返す
func (r *server) ListLabos(ctx context.Context, in *pbSchool.ListLabosRequest) (*pbSchool.ListLabosResponse, error) {
	pageSize := int32(5)
	if in.PageSize != nil {
		pageSize = *in.PageSize
	}
	offset := int32(0)
	if in.Offset != nil {
		offset = *in.Offset
	}
	results, err := r.Repo.ListLabos(ctx, pageSize, offset)
	if err != nil {
		return nil, err
	}
	var labos []*pbSchool.Labo
	for _, labo := range results {
		l, err := svcconv.LaboToProto(labo)
		if err != nil {
			return nil, err
		}
		labos = append(labos, l)
	}
	return &pbSchool.ListLabosResponse{Labos: labos, Offset: pageSize + offset}, err
}

func main() {
	// TCP ポートを作成する
	lis, err := net.Listen("tcp", ServerHost+":8080")
	if err != nil {
		log.Fatalln("TCP ポートのリッスンに失敗:", err)
	}

	// gRPC サーバーオブジェクトを作成する
	s := grpc.NewServer()
	// School サービスを接続
	pbSchool.RegisterShoolServer(s, &server{
		Repo: repository.NewSchool(infra.DB()),
	})
	// gRPC サーバーを起動
	log.Println("gRPC を起動 " + ServerHost + ":8080")
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
		Addr:    ServerHost + ":8090",
		Handler: gwmux,
	}

	log.Println("gRPC-Gateway 起動 http://" + ServerHost + ":8090")
	log.Fatalln(gwServer.ListenAndServe())
}
