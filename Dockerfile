FROM golang:1.21.5

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o poc .

EXPOSE 3428
EXPOSE 2574
EXPOSE 3526

CMD ["/app/poc", "mock"]