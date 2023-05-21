#!/bin/bash
# 第一个功能
cd ./service/app/user/rpc
go run user.go -f etc/user.yaml
# 第二个功能
cd ../api
go run user.go -f etc/user-api.yaml

