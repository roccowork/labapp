# ---------- 第 1 阶段：测试 + 编译 ----------
FROM golang:1.22-alpine AS builder
ENV CGO_ENABLED=0
WORKDIR /src
COPY go.mod ./
COPY *.go ./
# 测试不通过，镜像就构建失败
RUN go vet ./... && go test ./...
RUN go build -o /labapp .

# ---------- 第 2 阶段：运行 ----------
FROM alpine:3.20
RUN adduser -D -u 10001 app
COPY --from=builder /labapp /usr/local/bin/labapp
USER app
# 版本号由 Jenkins 构建时用 --build-arg 传入，比如 7-a1b2c3d
ARG APP_VERSION=dev
ENV APP_VERSION=${APP_VERSION}
EXPOSE 8000
CMD ["labapp"]