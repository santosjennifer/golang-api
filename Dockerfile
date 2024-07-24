FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o goApp ./cmd

# Imagem final
FROM scratch

COPY --from=builder /app/goApp /

EXPOSE 8000

ENTRYPOINT ["/goApp"]