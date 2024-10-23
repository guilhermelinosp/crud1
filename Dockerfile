FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o go-minikube-app

FROM alpine:latest

COPY --from=builder /app/go-minikube-app /go-minikube-app

EXPOSE 8080

CMD ["/go-minikube-app"]