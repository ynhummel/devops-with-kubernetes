FROM golang:1.26-alpine

ENV PORT=8001

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o todo-app .

EXPOSE $PORT

CMD ["./todo-app"]
