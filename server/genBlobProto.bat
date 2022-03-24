set PROTO_PATH=.\blob\api
set GO_OUT_PATH=.\blob\api\gen\v1
mkdir %GO_OUT_PATH%
protoc -I=%PROTO_PATH% --go_out=plugins=grpc,paths=source_relative:%GO_OUT_PATH% blob.proto