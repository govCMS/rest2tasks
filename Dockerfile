FROM golang as builder
RUN mkdir /app
COPY main.go /app
WORKDIR /app
RUN go build .

FROM alpine
WORKDIR /app
COPY --from=builder /app/app /app/rest2tasks
ENTRYPOINT ["/app/rest2tasks"]
