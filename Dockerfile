# 第一阶段：构建应用程序
FROM golang:1.21 AS build
WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=on && \
  go env -w GOPROXY=https://goproxy.cn,direct && \
  go build -o app

# 第二阶段：运行应用程序
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/app /app/app
RUN apk add --no-cache libc6-compat
COPY ./entrypoint.sh ./
ENTRYPOINT [ "sh", "entrypoint.sh" ]
