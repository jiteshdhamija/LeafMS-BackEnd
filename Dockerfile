FROM ubuntu:latest
LABEL authors="JITESH"
WORKDIR /DockerImages
RUN go get go.mongodb.org/mongo-driver/mongogo
RUN go get github.com/gorilla/mux

ENTRYPOINT ["go", "run", "main.go","db_connection.go"]