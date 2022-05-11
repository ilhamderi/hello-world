FROM golang:alpine
RUN apk update && apk add --no-cache git
RUN mkdir /app
RUN mkdir /test
COPY . /app
WORKDIR /app
RUN go mod init main
RUN go build -o main . 
CMD ["/app/main"]
