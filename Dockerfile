FROM golang:1.23

WORKDIR /usr/src/app
COPY go.* *.go ./
RUN go mod download && go mod verify

COPY modules modules/
RUN go build buy.go
EXPOSE 8080
CMD ["./buy"]