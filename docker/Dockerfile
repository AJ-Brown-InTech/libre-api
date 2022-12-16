#ubuntu 22.04 did not work with github workflows for creating a ci pipeline
# thats the only reason 20.04 is being used
FROM ubuntu:20.04 

#setup dir and fs
# RUN apt-get -y update
# RUN useradd -m admin


#add admin user and dir
RUN apt-get -y update
RUN useradd -m admin

# install golang v1.19 (doesn't matter what sub version)
RUN apt-get update
RUN apt-get install -y wget git gcc

RUN wget -P /tmp "https://dl.google.com/go/go1.19.4.linux-amd64.tar.gz"

RUN tar -C /usr/local -xzf "/tmp/go1.19.4.linux-amd64.tar.gz"
RUN rm "/tmp/go1.19.4.linux-amd64.tar.gz"

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
RUN go version

#create api dir
RUN cd home; cd admin;
RUN mkdir api

#run following cmd in specific api dir
WORKDIR /home/admin/api

#bundle app source
COPY ./ /home/admin/api
RUN go run main.go