FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY src/*.go ./

RUN go build -o /bittersweet-bot

# Pass this in with 'docker run -e BOT_TOKEN="your-token"'
ENV BOT_TOKEN=""

# sh is needed to expand the BOT_TOKEN ENV variable
CMD [ "sh", "-c", "/bittersweet-bot -t ${BOT_TOKEN}" ]
