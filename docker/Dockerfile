FROM golang:1.9.7-stretch
# RUN apt-get install -y mongodb
# ******Install mongo client**********
RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 9DA31620334BD75D9DCB49F368818C72E52529D4
RUN echo "deb http://repo.mongodb.org/apt/debian stretch/mongodb-org/4.0 main" | tee /etc/apt/sources.list.d/mongodb-org-4.0.list
RUN apt-get update && apt-get install -y mongodb-org
# ************************************

RUN go get -d -v github.com/gorilla/mux
RUN go install -v github.com/gorilla/mux
RUN go get -d -v gopkg.in/mgo.v2
RUN go install -v gopkg.in/mgo.v2
RUN go get -d -v github.com/codegangsta/gin
RUN go install -v github.com/codegangsta/gin
WORKDIR /go/src/app
COPY . .
