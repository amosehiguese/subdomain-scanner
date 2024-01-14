FROM golang:1.20.5-alpine AS build
WORKDIR /go/src/subd
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/subd

FROM alpine:3.18.4
RUN apk add --no-cache ca-certificates
WORKDIR /subd
COPY --from=build /go/bin/subd /bin/subd
COPY ./templates ./templates
COPY ./static ./static

EXPOSE 8080
ENTRYPOINT ["/bin/subd"]
