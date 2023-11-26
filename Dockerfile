FROM golang:1.21

WORKDIR /go/src/app

COPY . .

RUN go build -o serve cmd/server/main.go
RUN go build -o migrate cmd/migrate/main.go
RUN go build -o seed cmd/seeder/main.go

EXPOSE 8001

CMD ["./serve"]
