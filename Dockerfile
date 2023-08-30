FROM golang:1.20.5-alpine AS build
WORKDIR /go/src/subd
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/subd ./bin/subd

FROM scratch
COPY --from=build /go/bin/subd /bin/subd
ENTRYPOINT ["./bin/subd"]
