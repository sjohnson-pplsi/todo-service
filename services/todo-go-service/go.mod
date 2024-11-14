module github.com/sjohnson-pplsi/todo-service/services/todo-go-service

go 1.23.1

replace github.com/sjohnson-pplsi/todo-service/modules/todo-go-lib => ../../modules/todo-go-lib

replace github.com/cardboardrobots/esmongo => ../../modules/esmongo

require (
	github.com/cardboardrobots/baseerror v0.0.2
	github.com/cardboardrobots/config v0.0.0
	github.com/cardboardrobots/esmongo v0.0.0-00010101000000-000000000000
	github.com/cardboardrobots/listener v0.0.1
	github.com/google/uuid v1.6.0
	github.com/sjohnson-pplsi/todo-service/modules/todo-go-lib v0.0.0
	go.mongodb.org/mongo-driver v1.17.1
	google.golang.org/grpc v1.68.0
	google.golang.org/protobuf v1.35.1
)

require (
	github.com/golang/snappy v0.0.4 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/klauspost/compress v1.17.10 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
