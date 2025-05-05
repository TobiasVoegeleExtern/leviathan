# Stage 1: Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# --- KORRIGIERTER Build-Befehl ---
# Schreibe die Binary nach /app/server im Builder-Stage
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/server ./main.go

# Stage 2: Final Stage (minimales Image)
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /

# --- KORRIGIERTER Copy-Befehl ---
# Kopiere die Binary von /app/server aus dem Builder nach /server im finalen Image
COPY --from=builder /app/server .

EXPOSE 8080

# Specify the command to run when the container starts
ENTRYPOINT ["/server"]