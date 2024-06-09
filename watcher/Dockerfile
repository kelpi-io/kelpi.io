FROM golang:1.22

WORKDIR /usr/src/app
COPY . .
RUN go build -v -o /usr/local/bin/watcher .

CMD ["curl", "${RHOST}"]