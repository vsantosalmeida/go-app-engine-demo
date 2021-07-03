FROM golang:latest

WORKDIR WORKDIR /go/src/gitlab.olist.io/channels-integrations

ARG BIN_File=../../../bin/person-api.bin

RUN cp ${BIN_File} .

EXPOSE 8080

CMD ["./person-api"]