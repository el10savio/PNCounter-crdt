FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git 

RUN mkdir /pncounter

WORKDIR /pncounter

COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o /go/bin/pncounter


FROM scratch

COPY --from=builder /go/bin/pncounter /go/bin/pncounter

ENTRYPOINT ["/go/bin/pncounter"]

EXPOSE 8080