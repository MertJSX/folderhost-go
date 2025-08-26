FROM golang:1.25-alpine

WORKDIR /folderhost-go

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o app main.go

CMD [ "./app" ]

EXPOSE 5000