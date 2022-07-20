FROM golang:1.16 as build-env

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM scratch
COPY --from=build-env /go/bin/etisbew /etisbew
CMD ["/etisbew"]
