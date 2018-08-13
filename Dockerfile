FROM golang:1.9.1-alpine
RUN apk add --update git; \
    mkdir -p ${GOPATH}/api; \
    go get -u github.com/gorilla/mux
WORKDIR ${GOPATH}/api/.
COPY api/. ${GOPATH}/api/.
RUN go get github.com/halil9/GoRestAPI/api/models
RUN go get github.com/halil9/GoRestAPI/api/config
RUN go get gopkg.in/mgo.v2/bson
RUN go get github.com/halil9/GoRestAPI/api/DAO

EXPOSE 8000
ENTRYPOINT go run main.go