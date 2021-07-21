# 使用golang:1.16作为编译的基础镜像
FROM golang:1.16

# 容器环境变量添加，会覆盖默认的变量值
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

# 更改容器当地时间为上海时间
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# 设置容器的工作区
WORKDIR /build

# 将本地系统的go.mod和go.sum拷贝到容器
# 这里用go mod download的好处是下次构建镜像文件时，当go.mod和go.sum没有改变时，它是有缓存的，可以避免重复下载依赖包，加快构建。
COPY go.mod .
COPY go.sum .
RUN go mod download

# 把本地全部文件添加到容器的/build目录
COPY . .

# 编译：将cmd/server1/main.go编译成可执行的二进制文件server1，并保存在/build目录下
# cmd/server1/main.go、/build/cmd/server1/main.go、./cmd/server1/main.go都可以
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o server1 cmd/server1/main.go

# 仅用于说明服务暴露的端口号，运行容器的时候需要用参数-p来指定宿主机端口:服务暴露端口
EXPOSE 8081

# 启动服务
# ENTRYPOINT ["./server1"]、ENTRYPOINT ["/build/server1"]、CMD ["./server1"]、CMD ["/build/server1"]都可以
ENTRYPOINT ["./server1"]