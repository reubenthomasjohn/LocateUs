# Build stage
FROM golang:1.23.6-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o heatmap-api

# Final image 
FROM scratch
WORKDIR /root/
COPY --from=builder /app/heatmap-api .
COPY --from=builder /app/app.env .
CMD ["./heatmap-api"]
