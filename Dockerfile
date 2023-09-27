FROM golang:latest

WORKDIR /app

COPY . .

ENV CONFIG_PATH "./config/config.yaml"

EXPOSE 80

#download psql
RUN apt-get update 
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o main ./cmd/todos

CMD ["./main"]