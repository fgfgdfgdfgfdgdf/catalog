FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0


RUN apk update --no-cache && apk add --no-cache
WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /app/catalog ./cmd/catalog/


FROM scratch

WORKDIR /app
COPY --from=builder /app/catalog /app/catalog

COPY --from=builder /build/configs/main.yaml /app/configs/main.yaml

CMD ["./catalog"]
