FROM golang:1.7.4-alpine
MAINTAINER  sahilsk

ENV PORT=8080
EXPOSE ${PORT}
ENV SOURCE=/go/src/github.com/sahilsk/cloud-native-go
WORKDIR $SOURCE
RUN go install

ENTRYPOINT [ "cloud-native-go" ]