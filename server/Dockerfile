FROM golang:1.21.6

# RUN go build server
# ENV PORT 4000

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . .

RUN go build .
