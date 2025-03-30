FROM golang:1.22

WORKDIR /app
COPY . .
RUN go mod init go-app && \
go mod tidy && \
go build
CMD ["go-app"]
