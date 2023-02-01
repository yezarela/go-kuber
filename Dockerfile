FROM golang:1.19.4-alpine3.17 as builder

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o main main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/

EXPOSE 4001

CMD [ "/app/main" ]
