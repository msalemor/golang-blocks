# Build the application from source
FROM golang:alpine AS build-stage
#RUN apk update
#RUN apk add -U --no-cache ca-certificates && update-ca-certificates
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o backend

# Deploy the application binary into a lean image
FROM scratch AS final
WORKDIR /app

COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-stage /app/backend /app/backend
COPY --from=build-stage /app/public /app/public

ENV GIN_MODE=release
ENV PORT=8080

EXPOSE 8080
ENTRYPOINT ["/app/backend"]
