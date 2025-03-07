FROM golang:1.23.3 AS build

WORKDIR /usr/src/myapp
COPY ./ ./
RUN env CGO_ENABLED=0 GODEBUG=x509ignoreCN=0 go build -buildvcs=false -o weather ./cmd/server 

FROM build AS tester
RUN go test -v ./... || (echo "Tests failed" && exit 1)

FROM scratch
WORKDIR /
COPY --from=build /usr/src/myapp/weather ./

EXPOSE 8080
CMD [ "/weather" ]

