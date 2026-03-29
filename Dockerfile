FROM golang:1.26-alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /MarkdownToHTML .

FROM alpine:latest
WORKDIR /work
COPY --from=builder /MarkdownToHTML /usr/local/bin/MarkdownToHTML

ENTRYPOINT ["MarkdownToHTML"]