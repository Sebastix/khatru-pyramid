FROM golang:1.23 AS fetch
COPY go.mod go.sum /app
WORKDIR /app
RUN go mod download

FROM ghcr.io/a-h/templ:latest AS generate-stage
COPY --chown=65532:65532 . /app
RUN ["templ", "generate"]

FROM golang:1.23 AS builder
COPY --from=generate-stage /app /app
WORKDIR /app
RUN go build -o khatru-invite .

FROM ubuntu:latest
COPY --from=builder /app/khatru-invite /app/
ENV DATABASE_PATH="/app/db"
ENV USERDATA_PATH="/app/users.json"
CMD ["/app/khatru-invite"]

