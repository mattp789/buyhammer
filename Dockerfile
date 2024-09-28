FROM golang:1.23

WORKDIR /usr/src/app
RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - \
    && echo "deb http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list
RUN apt-get update && apt-get -y install google-chrome-stable
RUN chrome &
COPY go.* *.go ./
RUN go mod download && go mod verify

COPY modules modules/
RUN go build buy.go
EXPOSE 8080
CMD ["./buy"]