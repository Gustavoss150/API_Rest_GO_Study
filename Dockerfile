FROM golang:1.18

# set working directory
WORKDIR /go/src/app

# Copy the source code
COPY . .

EXPOSE 8000

RUN go build -o main cmd/main.go

CMD ["./main"]