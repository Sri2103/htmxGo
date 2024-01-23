# dockerfile for go?

FROM golang:1.21.6 as BUILDER

WORKDIR /app

COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . .

RUN go build -o myapp

# Use alpine to make the image smaller in size
FROM alpine:3.15 AS FINAL

WORKDIR /app
COPY --from=BUILDER /app/templates ./templates
COPY --from=BUILDER /app/static ./static

COPY --from=BUILDER /app/myapp  ./myapp

EXPOSE 3500

CMD ["./myapp"]