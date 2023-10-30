FROM golang AS builder
WORKDIR /app
RUN apt-get update -y && apt install protobuf-compiler -y
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
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
COPY cmd/server.crt /server.crt
COPY cmd/server.key /server.key
EXPOSE 4444
ENTRYPOINT ["/main"]