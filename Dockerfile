FROM golang AS build

WORKDIR /

COPY . .

RUN CGO_ENABLED=0 go build -o server ./cmd/main.go

FROM alpine

WORKDIR /app

COPY --from=build /server /app/server

EXPOSE 8080

ENTRYPOINT ["./server"]
