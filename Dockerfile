FROM dockerfile/go

RUN apt-get install -y curl git

# fluentd
RUN curl http://packages.treasuredata.com/GPG-KEY-td-agent | apt-key add -
RUN echo "deb http://packages.treasuredata.com/2/ubuntu/precise/ precise contrib" > /etc/apt/sources.list.d/treasure-data.list
RUN apt-get update
RUN apt-get install -y --force-yes td-agent
###

ENV GOPATH $HOME
ENV PATH $PATH:$GOPATH/bin
ADD ./src/main.go ./gopath/main.go
ADD ./src/1x1.gif ./gopath/1x1.gif
RUN go get github.com/t-k/fluent-logger-golang/fluent

EXPOSE 8080

RUN /usr/sbin/td-agent-gem install fluent-plugin-influxdb

ADD ./fluent.simple.conf /etc/td-agent/td-agent.conf.tmp
RUN cat /etc/td-agent/td-agent.conf.tmp > /etc/td-agent/td-agent.conf

ENTRYPOINT service td-agent restart \
	&& go run main.go
