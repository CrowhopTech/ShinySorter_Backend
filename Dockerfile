# Can't use the minimal/slim golang image as it doesn't have "make"
FROM golang:1.18

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN make bin/dbpopulator bin/restserver

# Using debian-slim instead of alpine as alpine has gcc issues
FROM debian:stable-slim
COPY --from=0 /app/bin/* /
