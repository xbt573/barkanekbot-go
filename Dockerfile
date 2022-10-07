FROM golang:latest AS build

WORKDIR /app
COPY . .

RUN go build

FROM alpine:latest

RUN apk add libc6-compat

WORKDIR /app
COPY --from=build /app/barkanekbot-go .

CMD /app/barkanekbot-go
