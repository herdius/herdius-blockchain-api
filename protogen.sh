#!/bin/bash

protoc --go_out=plugins=grpc,import_path=core_grpc:. protobuf/message.proto