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