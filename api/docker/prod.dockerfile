FROM golang:1.14 AS build

WORKDIR /home/dsdit

COPY api ./api/
COPY go.mod go.sum ./

RUN CGO_ENABLED=0 go build -o bin/server ./api/*.go 

FROM alpine:latest
WORKDIR /home/dsdit
COPY --from=build /home/dsdit/bin/server .
ENTRYPOINT [ "./server" ]