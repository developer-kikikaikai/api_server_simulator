#base image
FROM ubuntu:18.04

ENV DEBCONF_NOWARNINGS yes

# user
LABEL maintainer="developer-kikikaikai <kikikaikaikai@gmail.com>"

# RUN: action to build docker
RUN echo "## Setup base tools"
RUN apt update
RUN apt install -y iproute2 iputils-ping curl
RUN echo "## Install golang"
RUN apt -y install golang
RUN apt -y install git
RUN echo "## Install mongo"
RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 0C49F3730359A14518585931BC711F9BA15703C6;
RUN echo "deb [ arch=amd64,arm64 ] http://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/3.4 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-3.4.list
RUN apt update
RUN apt -y install mongodb-org
RUN mkdir -p /data/db
RUN echo "## Install deps"
RUN apt -y install go-dep
RUN echo "## get code"
RUN go get -v github.com/developer-kikikaikai/api_server_simulator
RUN cd /root/go/src/github.com/developer-kikikaikai/api_server_simulator && make
RUN chmod +x /root/go/src/github.com/developer-kikikaikai/api_server_simulator/docker/startServices.sh

# execute command and port to run docker
ENTRYPOINT /root/go/src/github.com/developer-kikikaikai/api_server_simulator/docker/startServices.sh
