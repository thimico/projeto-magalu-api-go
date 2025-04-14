FROM golang:1.23-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go get golang.org/x/sys@v0.30.0

RUN go install github.com/githubnemo/CompileDaemon@latest

COPY . .

RUN go get github.com/githubnemo/CompileDaemon
ENV DATABASE_URL mongodb://mongo:27017/
ENV DATABASE_NAME magalu-api
ENV TELEGRAM_BOT_TOKEN token
ENV TELEGRAM_CHAT_ID chat
ENV PORT 8080

EXPOSE 8080
WORKDIR /app

ENTRYPOINT ["CompileDaemon", "-exclude-dir=.git", "-exclude-dir=docs", "--build=go build ./cmd/main.go", "--command=./main"]
