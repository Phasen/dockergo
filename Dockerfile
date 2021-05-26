FROM golang:alpine
RUN apk --no-cache add ca-certificates
WORKDIR /go/src/github/Phasen/billymulrine/
COPY src .
CMD ["go","run","./main.go"] 