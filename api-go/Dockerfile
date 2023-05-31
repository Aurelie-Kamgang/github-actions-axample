FROM golang:latest
COPY . /app
WORKDIR /app
RUN go build -o app
EXPOSE 8080
CMD ["./app"]
