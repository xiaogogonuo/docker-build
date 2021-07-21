# 使用golang:1.16作为编译的基础镜像
FROM golang:1.16 AS build
# 容器环境变量添加，会覆盖默认的变量值
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
# 设置容器build阶段的工作区
WORKDIR /go/release

COPY go.mod .
COPY go.sum .
RUN go mod download
# 把全部文件添加到/go/release目录
ADD . .
# 编译：将cmd/server2/main.go编译成可执行的二进制文件server2，并保存在/go/release目录下
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o server2 cmd/server2/main.go


# 运行：使用alpine作为基础镜像
FROM alpine AS run
# 设置容器run阶段的工作区
WORKDIR /deploy
# 在build阶段复制时区到/etc/localtime
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# 将build阶段的可执行文件server2复制到run阶段的工作区
COPY --from=build /go/release/server2 /
# 将build阶段的配置文件复制到run阶段的工作区
COPY --from=build /go/release/configs configs

# 启动服务
CMD ["/server2"]