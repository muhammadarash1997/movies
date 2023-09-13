FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go build -o binary-movies

EXPOSE 8080

CMD ./binary-movies