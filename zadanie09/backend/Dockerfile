FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/server .

FROM alpine:3.20
WORKDIR /app
COPY --from=build /out/server /app/server
EXPOSE 8080
CMD ["/app/server"]
