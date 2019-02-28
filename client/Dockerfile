FROM golang:alpine AS builder

LABEL maintainer="Gangadhar Pathipaka"

RUN apk update && apk add --no-cache git

WORKDIR /go/src/go-docker/client

COPY . .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /go/bin/client

FROM scratch

COPY --from=builder /go/bin/client /go/bin/client

ENTRYPOINT [ "/go/bin/client" ]
