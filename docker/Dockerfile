FROM golang:alpine

ENV BIN_FILE /server
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build \
        -o ${BIN_FILE} ./cmd/web/*

EXPOSE 8080
EXPOSE 27017
EXPOSE 6379

CMD ${BIN_FILE}