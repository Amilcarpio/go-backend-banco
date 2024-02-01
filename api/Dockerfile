FROM golang:1.21.5

ENV GOPLUGIN_MODE=release
ENV PORT 3306

WORKDIR /go/src/app

COPY . .

RUN go mod download
RUN go build -o /go/bin/app cmd/main.go

CMD ["/go/bin/app"]