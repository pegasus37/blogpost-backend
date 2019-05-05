#===========
#Build Stage
#===========
FROM golang:1.12-alpine as builder

# install basic dependencies
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates bash bash-completion tzdata

# copy app folder to gopath
ENV GOPATH /code
ENV PATH /code/bin:/usr/local/go/bin:$PATH

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/msoap/go-carpet

RUN mkdir -p /code/src/github.com/pegasus37/blogpost-backend

FROM alpine:3.7
# Set your app in '/app' directory and set it as current directory
WORKDIR /app/

EXPOSE 8081

 # Run
CMD [""]

