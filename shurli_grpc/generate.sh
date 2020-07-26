#!/bin/bash

export GO111MODULE=on

protoc shurlipb/shurli.proto --go_out=plugins=grpc:.
protoc -I shurlipb/ shurlipb/shurli.proto --dart_out=grpc:shurlipb/dart