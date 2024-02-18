FROM golang:1.21-alpine AS build
WORKDIR /app
COPY api/go.* .
ENV GO111MODULE=on
RUN go mod download
COPY api/. .
RUN CGO_ENABLED=0 GOOS=linux go build -o api cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/api ./
ENTRYPOINT ["./api"]