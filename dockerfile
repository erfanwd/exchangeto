FROM golang:1.19.3-alpine3.16 as builder

WORKDIR /app
COPY /. ./

#apk add nano \ && 
RUN go mod download \
        && go build -o /github.com/erfanwd/exchangeto .

FROM alpine:3.16.2
WORKDIR /app
COPY --from=builder github.com/erfanwd/exchangeto .
COPY .env .


CMD ["./github.com/erfanwd/exchangeto"]
