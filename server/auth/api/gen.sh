# 由Proto文件生成对应的Go/GRPC/GRPC-Gateway代码
GO_OUT_PATH=gen/v1
mkdir -p $GO_OUT_PATH
PROTO_FILE=auth.proto
YAML_FILE=auth.yaml
protoc -I=. --go_out=paths=source_relative:$GO_OUT_PATH $PROTO_FILE 
protoc -I=. --go-grpc_out=paths=source_relative:$GO_OUT_PATH $PROTO_FILE
protoc -I=. --grpc-gateway_out=paths=source_relative,grpc_api_configuration=$YAML_FILE:$GO_OUT_PATH $PROTO_FILE

# 由Proto文件生成前端的JS/TS代码
PBTS_OUT_DIR=../../../wx/miniprogram/service/proto_gen/auth
JS_FILE=auth_pb.js
TS_FILE=auth_pb.d.ts
mkdir -p $PBTS_OUT_DIR
pbjs -t static -w es6 $PROTO_FILE \
--no-create --no-encode --no-decode \
--no-verify --no-delimited --force-number \
-o $PBTS_OUT_DIR/tmp.js 

echo 'import * as $protobuf from "protobufjs";\n' > $PBTS_OUT_DIR/$JS_FILE
cat $PBTS_OUT_DIR/tmp.js >> $PBTS_OUT_DIR/$JS_FILE
rm $PBTS_OUT_DIR/tmp.js
pbts -o $PBTS_OUT_DIR/$TS_FILE $PBTS_OUT_DIR/$JS_FILE