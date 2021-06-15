# 构建基镜像
FROM golang:latest as builder
MAINTAINER Kaneziki<1848224883@qq.com>
# 拷贝项目文件
COPY ./ $GOPATH/src/user/
WORKDIR $GOPATH/src/user/
# 构建基础环境
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 GOOS=linux go build -o user main.go

# 构建二级镜像
FROM alpine as prod
COPY --from=builder /go/src/user/ /user/
WORKDIR /user/
# 暴露端口
EXPOSE 8000
ENTRYPOINT ./user --consul.host=$consulHost --service.address=$serviceAddr