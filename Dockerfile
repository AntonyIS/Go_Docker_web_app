FROM golang:latest

RUN mkdir /build

WORKDIR /build

RUN export GO111MODULE=on

RUN go get github.com/AntonyIS/Go_Docker_web_app/main

RUN cd /build && git clone https://github.com/AntonyIS/Go_Docker_web_app.git

RUN cd /build/Go_Docker_web_app/main && go build

EXPOSE 5000

ENTRYPOINT ["/build/Go_Docker_web_app/main/main"]