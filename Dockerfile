FROM golang:1.18-buster AS build

WORKDIR /app

COPY ./ ./
RUN go mod download
RUN go build server.go

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/server /server
COPY --from=build /app/dev.db /dev.db

EXPOSE 8080

ENTRYPOINT ["/server"]