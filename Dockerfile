FROM golang:alpine as build

RUN apk add build-base
RUN export GOBIN=$GOPATH/bin
WORKDIR /go/src/strider
COPY . ./

RUN go mod download
RUN go build -o bin/main main.go
RUN ls

FROM golang:alpine
COPY --from=build /go/src/strider/bin/main main
ENTRYPOINT [ "./main" ]
EXPOSE 8080

