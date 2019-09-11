FROM golang:alpine
LABEL maintainer="Elias Krontiris <teliaz@gmail.com>"
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o goapi .
EXPOSE 8080
CMD ["./goapi"]


