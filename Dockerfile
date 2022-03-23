FROM golang:1.17 as build

WORKDIR /app
COPY . . 
RUN CGO_ENABLED=0 go build -o server main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/server .
CMD ["./server"]
