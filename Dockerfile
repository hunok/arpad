FROM golang:1.15-alpine AS build
WORKDIR /src
COPY . .

RUN apk update && apk add git ca-certificates build-base

RUN go get -d -v

RUN go test -cover -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-w -s" -o bin/arpad .

FROM scratch

LABEL "maintainer"="hunok/arpad <arpad@akos.me>"
LABEL "version"="0.1"
LABEL "description"="github.com/hunok/arpad"

COPY --from=build /src/bin/arpad /bin/arpad

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["/bin/arpad"]
