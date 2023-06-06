FROM golang:1.19-alpine

WORKDIR /notesApp
RUN apk update && apk add --no-cache git gcc musl-dev
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN go build -o main .

EXPOSE 8080:8080

CMD ["./main"]
