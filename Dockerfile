FROM ubuntu:18.04

RUN apt-get update
RUN apt-get install -y ca-certificates
RUN mkdir /app

ADD main /app

WORKDIR /app

EXPOSE 8080

CMD ["./main"]