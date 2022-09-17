FROM golang:alpine AS Builder
WORKDIR /app/bing-wallpaper/

RUN apk add build-base

COPY . .
RUN go mod download
RUN go build -o server .

FROM alpine AS Runner
WORKDIR /app/bing-wallpaper/

ENV LISTEN :8080
EXPOSE 8080

COPY --from=Builder /app/bing-wallpaper/server server

RUN chmod +x server
CMD ./server
