FROM golang:latest AS build

WORKDIR /src/

COPY main.go /src/

RUN CGO_ENABLED=0 go build -o /bin/demo

FROM scratch
COPY --from=build /bin/demo /bin/demo
ENTRYPOINT ["/bin/demo"]
