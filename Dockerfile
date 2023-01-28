FROM golang:1.19.5-alpine3.17

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o /godocker
EXPOSE 8080

CMD [ "/godocker" ]