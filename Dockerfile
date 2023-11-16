# 第一阶段：构建应用程序
FROM golang:latest AS build
WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=on && \
  go env -w GOPROXY=https://goproxy.cn,direct && \
  go build -o app

# 第二阶段：运行应用程序
FROM alpine:latest
WORKDIR /app
COPY --from=build /app /app
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY ./entrypoint.sh ./
ENTRYPOINT [ "sh","entrypoint.sh" ]
