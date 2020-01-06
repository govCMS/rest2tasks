FROM golang:alpine as builder
RUN mkdir /app
COPY main.go /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/rest2tasks

FROM scratch
COPY --from=builder /bin/rest2tasks /bin/rest2tasks
ENTRYPOINT ["/bin/rest2tasks"]
