FROM golang:latest
WORKDIR /app

COPY . .
RUN go env -w GO111MODULE=auto
RUN go build -o ./auth-service/main ./auth-service/main.go
WORKDIR /app/auth-service
ENTRYPOINT ["./main"]

FROM golang:latest
WORKDIR /app

COPY . .
RUN go env -w GO111MODULE=auto
RUN go build -o ./gateway-service/main ./gateway-service/main.go
WORKDIR /app/gateway-service
ENTRYPOINT ["./main"]

# 构建log镜像
FROM golang:latest
WORKDIR /app

COPY . .
RUN go env -w GO111MODULE=auto
RUN go build -o ./log-service/main ./log-service/main.go
WORKDIR /app/log-service
ENTRYPOINT ["./main"]

FROM golang:latest
WORKDIR /app

COPY . .
RUN go env -w GO111MODULE=auto
RUN go build -o ./store-service/main ./store-service/main.go
WORKDIR /app/store-service
ENTRYPOINT ["./main"]

FROM golang:latest
WORKDIR /app

COPY . .
RUN go env -w GO111MODULE=auto
RUN go build -o ./task-service/main ./task-service/main.go
WORKDIR /app/task-service
ENTRYPOINT ["./main"]

# 底层数据库微服务
FROM golang:latest
WORKDIR /app

COPY . .
RUN go env -w GO111MODULE=auto
RUN go build -o ./store-service-grpc/mongodb-first/main ./store-service-grpc/mongodb-first/main.go
WORKDIR /app/store-service-grpc/mongodb-first
ENTRYPOINT ["./main"]

FROM golang:latest
WORKDIR /app

COPY . .
RUN go env -w GO111MODULE=auto
RUN go build -o ./store-service-grpc/mysql-first/main ./store-service-grpc/mysql-first/main.go
WORKDIR /app/store-service-grpc/mysql-first
ENTRYPOINT ["./main"]

FROM golang:latest
WORKDIR /app

COPY . .
RUN go env -w GO111MODULE=auto
RUN go build -o ./store-service-grpc/mysql-second/main ./store-service-grpc/mysql-second/main.go
WORKDIR /app/store-service-grpc/mysql-second
ENTRYPOINT ["./main"]