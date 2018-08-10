FROM nimmis/golang:latest
COPY api/. /api/.
WORKDIR /api/.
RUN go get "github.com/gorilla/mux"
RUN go get "github.com/halil9/GoRestAPI/config"
RUN go get "github.com/halil9/GoRestAPI/DAO"
RUN go get "github.com/halil9/GoRestAPI/models"
RUN go get "gopkg.in/mgo.v2/bson"
EXPOSE 8000
ENTRYPOINT go run main.go