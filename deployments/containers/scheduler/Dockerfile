FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/

RUN go build -C cmd/scheduler -o /oci-resource-manager

CMD [ "/oci-resource-manager" ]