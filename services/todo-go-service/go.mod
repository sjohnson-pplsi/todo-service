module github.com/sjohnson-pplsi/todo-service/services/todo-go-service

go 1.23.1

replace github.com/sjohnson-pplsi/todo-service/modules/todo-go-lib => ../../modules/todo-go-lib

require (
	github.com/cardboardrobots/baseerror v0.0.2
	github.com/cardboardrobots/config v0.0.0
	github.com/cardboardrobots/listener v0.0.1
	github.com/sjohnson-pplsi/todo-service/modules/todo-go-lib v0.0.0
	google.golang.org/grpc v1.68.0
)

require (
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
