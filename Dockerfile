FROM golang:1.23.3 AS build
WORKDIR /usr/src/myapp
COPY ./ ./
RUN env CGO_ENABLED=0 go build -o weather -buildvcs=false

FROM scratch
WORKDIR /
COPY --from=build /usr/src/myapp/weather ./
EXPOSE 8080
CMD [ "/weather" ]
