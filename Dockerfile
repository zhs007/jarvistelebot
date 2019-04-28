FROM golang:1.12 as builder

MAINTAINER zerro "zerrozhao@gmail.com"

WORKDIR /src/jarvistelebot

COPY ./go.* /src/jarvistelebot/

RUN go mod download

COPY . /src/jarvistelebot

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jarvistelebot . \
    && mkdir /app \
    && mkdir /app/jarvistelebot \
    && mkdir /app/jarvistelebot/cfg \
    && mkdir /app/jarvistelebot/dat \
    && mkdir /app/jarvistelebot/logs \
    && cp ./jarvistelebot /app/jarvistelebot/ \
    && cp -r www /app/jarvistelebot/www \
    && cp ./cfg/config.yaml.default /app/jarvistelebot/cfg/config.yaml

FROM alpine
RUN apk upgrade && apk add --no-cache ca-certificates
WORKDIR /app/jarvistelebot
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /app/jarvistelebot /app/jarvistelebot
CMD ["./jarvistelebot"]
