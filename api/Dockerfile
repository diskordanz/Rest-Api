FROM golang:1.9

RUN mkdir -p /apiLib
WORKDIR /apiLib
COPY ./library .
EXPOSE 8080
ENTRYPOINT ["./library"]

# RUN mkdir -p /go/src/rest-api/api
# WORKDIR /go/src/rest-api/api
# COPY . /go/src/rest-api/api
# RUN go-wrapper download
# RUN go-wrapper install
# EXPOSE 8080
# CMD ["/go/src/rest-api/api", "run", "main.go"]