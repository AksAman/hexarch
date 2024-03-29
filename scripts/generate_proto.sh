# !/usr/bin/bash

grpc_dir="internal/adapters/framework/left/grpc"
src_proto_dir="$grpc_dir/proto"


protoc --go_out=$grpc_dir --proto_path=$src_proto_dir $src_proto_dir/*_msg.proto
protoc --go-grpc_out=require_unimplemented_servers=false:$grpc_dir --proto_path=$src_proto_dir $src_proto_dir/*_service.proto
