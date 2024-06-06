FROM golang:1.22.1-alpine

WORKDIR /
RUN apk add make

COPY go.mod ./
RUN go mod download

COPY . ./

RUN make build

EXPOSE 8080

CMD [ "/main" ]