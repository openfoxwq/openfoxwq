rm -f proto/*_foxwqrpc.pb.go

go install github.com/ale64bit/openfoxwq/cmd/protoc-gen-go-foxwqrpc && protoc \
   --cpp_out=../openfoxwq_qtclient \
   --go_out=. \
   --go-foxwqrpc_out=. \
   --go_opt=paths=source_relative \
   --go-foxwqrpc_opt=paths=source_relative \
   proto/reqOption.proto \
   proto/common.proto \
   proto/fox.proto \
   proto/nav.proto \
   proto/broadcast.proto \
   proto/play.proto \
   proto/ws.proto
