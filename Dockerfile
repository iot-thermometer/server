FROM golang AS builder
WORKDIR /app
RUN apt-get update -y && apt install protobuf-compiler -y
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN git submodule update --init
RUN make
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /main github.com/iot-thermometer/server/cmd

FROM alpine
WORKDIR /
COPY --from=builder /main /main
EXPOSE 4444
ENTRYPOINT ["/main"]