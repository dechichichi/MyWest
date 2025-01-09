# 使用Go 1.23官方镜像作为基础镜像
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

# 将当前目录下的所有文件复制到工作目录中
COPY . .

# 下载依赖项并编译应用程序
RUN go mod download
RUN go build -o todolist

# 使用一个更小的基础镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /root/

# 从builder阶段复制编译后的二进制文件到当前镜像
COPY --from=builder /app/todolist .

# 暴露应用程序运行的端口
EXPOSE 8080

# 启动应用程序
CMD ["./todolist"]