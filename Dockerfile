# ---------- Builder ----------
FROM golang:1.23.3 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /mailq .

# ---------- Runtime ----------
FROM scratch

COPY --from=builder /mailq /mailq

ENTRYPOINT ["/mailq"]
