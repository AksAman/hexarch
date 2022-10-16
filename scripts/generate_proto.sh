# !/usr/bin/bash

grpc_dir="internal/adpaters/framework/left/grpc"
src_proto_dir="$grpc_dir/proto"


protoc --go_out=$grpc_dir --proto_path=$src_proto_dir $src_proto_dir/*_msg.proto
protoc --go-grpc_out=$grpc_dir --proto_path=$src_proto_dir $src_proto_dir/*_service.proto
