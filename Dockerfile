FROM golang:1.10 as builder

MAINTAINER zerro "zerrozhao@gmail.com"

WORKDIR $GOPATH/src/github.com/zhs007/jarvistelebot

COPY ./Gopkg.* $GOPATH/src/github.com/zhs007/jarvistelebot/

RUN go get -u github.com/golang/dep/cmd/dep \
    && dep ensure -vendor-only -v

COPY . $GOPATH/src/github.com/zhs007/jarvistelebot

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jarvistelebot . \
    && mkdir /home/jarvistelebot \
    && mkdir /home/jarvistelebot/cfg \
    && mkdir /home/jarvistelebot/dat \
    && mkdir /home/jarvistelebot/logs \
    && cp ./jarvistelebot /home/jarvistelebot/ \
    && cp ./cfg/config.yaml.default /home/jarvistelebot/cfg/config.yaml

FROM alpine
RUN apk upgrade && apk add --no-cache ca-certificates
WORKDIR /home/jarvistelebot
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /home/jarvistelebot /home/jarvistelebot
CMD ["./jarvistelebot"]
