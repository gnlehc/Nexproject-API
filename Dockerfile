FROM golang:1.22-alpine AS build

WORKDIR /app
COPY . .

RUN go build -o main

FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/main .
COPY .env . 
EXPOSE 8080

CMD ["./main"]