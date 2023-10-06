FROM golang AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /main github.com/iot-thermometer/server/cmd

FROM alpine
WORKDIR /
COPY --from=builder /main /main
EXPOSE 4444
ENTRYPOINT ["/main"]