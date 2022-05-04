set PROTO_PATH=.\car\api
set GO_OUT_PATH=.\car\api\gen\v1
mkdir %GO_OUT_PATH%
protoc -I=%PROTO_PATH% --go_out=plugins=grpc,paths=source_relative:%GO_OUT_PATH% car.proto
protoc -I=%PROTO_PATH% --grpc-gateway_out=paths=source_relative,grpc_api_configuration=%PROTO_PATH%/car.yaml:%GO_OUT_PATH% car.proto

set PBTS_BIN_DIR= ..\wx\miniprogram\node_modules\.bin
set PBTS_OUT_DIR= ..\wx\miniprogram\service\proto_gen\car
mkdir %PBTS_OUT_DIR%

%PBTS_BIN_DIR%\pbjs -t static -w es6 ./car/api/car.proto --no--create --no--decode --no--verify --force-number --no--delimited -o %PBTS_OUT_DIR%/car_pb.js

%PBTS_BIN_DIR%\pbts -o %PBTS_OUT_DIR%/car_pb.d.ts %PBTS_OUT_DIR%/car_pb.js