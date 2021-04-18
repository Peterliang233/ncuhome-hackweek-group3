FROM golang:apline

ENV GO111MODULE=on\
    CGO_ENABLED=0\
    COOS=linux\
    GOARCH=amd64\
    GOPROXY=https://goproxy.cn,direct

WORKDIR /build

COPY . .

RUN go build -o app.

WORKDIR /dist

RUN cp /build/app.

EXPOSE 80

CMD  ["/dist/app"]