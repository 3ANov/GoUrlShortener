FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .
RUN go build -o /build/main main.go

FROM alpine

WORKDIR /backend

COPY --from=builder /build/main /backend/main
COPY ./config.json /backend/config.json

CMD ["/backend/main"]
