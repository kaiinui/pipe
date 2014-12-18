FROM ubuntu:12.04

RUN apt-get update && apt-get -y install curl git hg golang

# fluentd
RUN curl -O http://packages.treasure-data.com/debian/RPM-GPG-KEY-td-agent && apt-key add RPM-GPG-KEY-td-agent && rm RPM-GPG-KEY-td-agent
RUN curl -L http://toolbelt.treasuredata.com/sh/install-ubuntu-precise.sh | sh

ENV GOPATH $HOME
ENV PATH $PATH:$GOPATH/bin
ADD ./src/main.go main.go
ADD ./src/1x1.gif 1x1.gif
RUN go build main.go

EXPOSE 3000

CMD ["./main"]

