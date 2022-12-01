FROM golang

WORKDIR /app

COPY go.mod /app
COPY go.sum /app
RUN go mod download

COPY . /app

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]