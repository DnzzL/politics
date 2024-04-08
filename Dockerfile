FROM golang:1.21-bookworm as builder

WORKDIR /app

RUN apt-get update \
 && DEBIAN_FRONTEND=noninteractive \
    apt-get install --no-install-recommends --assume-yes \
      build-essential \
      libsqlite3-dev

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/a-h/templ/cmd/templ@latest \
    && go install github.com/go-jet/jet/v2/cmd/jet@latest

COPY . .

RUN make init-db \
    && chmod a+rw politics.sqlite

RUN CGO_ENABLED=1 GOOS=linux make build


FROM debian:bookworm

WORKDIR /app

COPY --from=builder /app .

ENV PORT=8080

EXPOSE $PORT

CMD ["./bin/main"]