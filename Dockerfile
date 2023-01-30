FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /simple-app main.go

EXPOSE 4001

CMD [ "/simple-app" ]
