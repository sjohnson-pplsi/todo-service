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
	"github.com/sjohnson-pplsi/todo-service/services/todo-go-service/feature/todo/infrastructure"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	log.SetPrefix("[liara]\t")
	log.Println("started...")

	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(conf.MongoDbUri))
	if err != nil {
		log.Fatal(err)
	}

	listener.Listen(context.Background(), conf.Port, conf.Port+1,
		http.NewServeMux(),
		initService(mongoClient))
	fmt.Println("running...")
}

func initService(mongoClient *mongo.Client) *grpc.Server {
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

	db := mongoClient.Database("todo-service")

	pb.RegisterTodoServiceServer(s,
		controller.NewTodoController(
			service.NewTodoService(
				infrastructure.NewTodoRepository(
					db.Collection("todo")))))

	return s
}
