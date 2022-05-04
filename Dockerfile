FROM golang:1.15 AS builder
COPY . /go/src/github.com/adefemi171/appdemo
WORKDIR /go/src/github.com/adefemi171/appdemo
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app main.go

FROM scratch
COPY --from=builder /app .
EXPOSE 3000
CMD ["/app"]