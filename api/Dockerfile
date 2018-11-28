FROM golang:1.9

RUN mkdir -p /apiLib
WORKDIR /apiLib
COPY ./library .
EXPOSE 8080
ENTRYPOINT ["./library"]




