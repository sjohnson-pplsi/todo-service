package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/cardboardrobots/listener"
	pb "github.com/sjohnson-pplsi/todo-service/modules/todo-go-lib/generated"
	"github.com/sjohnson-pplsi/todo-service/services/todo-go-service/config"
	"github.com/sjohnson-pplsi/todo-service/services/todo-go-service/feature/base"
	"github.com/sjohnson-pplsi/todo-service/services/todo-go-service/feature/todo/controller"
	"github.com/sjohnson-pplsi/todo-service/services/todo-go-service/feature/todo/domain/service"
	"google.golang.org/grpc"
)

func main() {
	log.SetPrefix("[liara]\t")
	log.Println("started...")

	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	listener.Listen(context.Background(), conf.Port, conf.Port+1,
		http.NewServeMux(),
		initService())
	fmt.Println("running...")
}

func initService() *grpc.Server {
	s := listener.NewServerBuilder().
		AddUnary(
			listener.LogGRPC(false),
			listener.ErrorInterceptor(base.GetStatusCodeGRPC),
		).
		AddStream(
			listener.LogStreamGRPC(false),
			listener.ErrorInterceptorStream(base.GetStatusCodeGRPC),
		).
		Build()

	// ctx := context.Background()

	pb.RegisterTodoServiceServer(s,
		controller.NewTodoController(
			service.NewTodoService(nil)))

	return s
}
