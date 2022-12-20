#!/usr/bin/env just --justfile
UPX_HOME := '/home/andrew/Applications/upx/'
run:
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app
  rm -f app.upx
  {{UPX_HOME}}/upx --lzma --best app -o app.upx
  cp ./app.upx ./infra/app
  docker build -t fasthttp-golang:1 ./infra
  docker run -d -it --memory=100m --cpus=4 -p 8089:8089 --network fast-http fasthttp-golang:1