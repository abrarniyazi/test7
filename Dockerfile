FROM golang:alpine
ENV CGO_ENABLED=0
WORKDIR /ACCOUNTAPI

COPY . .

RUN go build -o ACCOUNTAPI

RUN echo "hello there"

RUN go test