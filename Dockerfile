FROM golang

WORKDIR /go/src/app

COPY . .

RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "main.go"]