files = \
	-Iprotos google/protobuf/timestamp.proto \
	-Iprotos protos/todo.proto

dart_out = ../modules/todo-dart-lib/lib/generated
go_out = ../modules/todo-go-lib/generated
ts_out = ../modules/todo-ts-lib/src/generated

.PHONY: default
default: dart go ts

.PHONY: dart
dart:
	rm -rf ${dart_out}
	mkdir -p ${dart_out}
	protoc \
	--dart_out=grpc:${dart_out} \
	$(files)

.PHONY: go
go:
	rm -rf ${go_out}
	mkdir -p ${go_out}
	protoc \
	--go_out=${go_out} \
	--go_opt=paths=source_relative \
	--go-grpc_out=${go_out} \
	--go-grpc_opt=paths=source_relative \
	$(files)

.PHONY: ts
ts:
	rm -rf ${ts_out}
	mkdir -p ${ts_out}
	grpc_tools_node_protoc \
	--js_out=import_style=commonjs,binary:${ts_out} \
	--grpc_out=grpc_js:${ts_out} \
	--ts_out=grpc_js:${ts_out} \
	$(files)

.PHONY: install
install: install-compilers install-go install-dart install-ts

.PHONY: install-compilers
install-compilers:
	brew install protobuf
	brew install clang-format

.PHONY: install-go
install-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

.PHONY: install-dart
install-dart:
	dart pub global activate protoc_plugin
	# export PATH="$PATH:$HOME/.pub-cache/bin"

.PHONY: install-ts
install-ts:
	npm install -g grpc-tools
	npm install -g grpc_tools_node_protoc_ts

# brew install protoc-gen-grpc-web
# npm install -g protoc-gen-js
# npm install -g protoc-gen-ts
