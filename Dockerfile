FROM golang:1.20

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait_for_postgres.sh

RUN go mod download
RUN go build -o avito-app ./cmd/main.go

CMD ["./avito-app"]