FROM golang:1.14.2

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN make build
EXPOSE 9000
CMD ["./ms-api"]