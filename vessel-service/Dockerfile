FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /go/src/go-docker/vessesl-service

COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /go/bin/vessel-service

FROM scratch

COPY --from=builder /go/bin/vessel-service /go/bin/vessel-service

EXPOSE 8081

ENTRYPOINT [ "/go/bin/vessel-service" ]