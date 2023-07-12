#!/bin/bash

protoc shopping.proto --go_out=../service --go-grpc_out=../service