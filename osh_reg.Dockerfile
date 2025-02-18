FROM golang:1.23.2-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /registrar ./cmd/main.go

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /registrar /registrar
# COPY config/registrarConfig.yaml /osh/configs/registrarConfig.yaml

EXPOSE 8070

ENTRYPOINT ["/registrar", "-config", "osh/configs/registrarConfig.yaml"]