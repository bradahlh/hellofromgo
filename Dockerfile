FROM golang:latest AS build

WORKDIR /src/

COPY main.go go.* /src/

RUN CGO_ENABLED=0 go build -o /bin/hellofromgo

FROM scratch
COPY --from=build /bin/hellofromgo /bin/hellofromgo
ENTRYPOINT ["/bin/hellofromgo"]
