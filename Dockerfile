# #default os but development can be in any os
# FROM ubuntu:22.04 

# #add admin user and dir
# RUN apt-get -y update
# RUN useradd -m admin && apt-get install -y fish && apt-get install -y golang-go && apt-get install -y ca-certificates
# #create api dir
# RUN cd home && cd admin && mkdir api && cd api 
# # #bundle app source
# COPY . ./home/admin/api

# #ru specific api dir
# WORKDIR /home/admin/api

# RUN go get
# RUN su - admin
# CMD ["echo","complete"]
# EXPOSE 8080

FROM ubuntu:latest

RUN apt-get update
RUN apt-get install -y wget git gcc
RUN apt-get update && apt-get install --no-install-recommends --yes python3
RUN useradd -m admin && apt-get install -y fish && apt-get install -y ca-certificates
RUN wget -P /tmp "https://dl.google.com/go/go1.18.linux-amd64.tar.gz"
RUN tar -C /usr/local -xzf "/tmp/go1.18.linux-amd64.tar.gz"
RUN rm "/tmp/go1.18.linux-amd64.tar.gz"
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

#bundle app source
COPY . ./home/admin/api
RUN cd home && cd admin && cd api 
WORKDIR /home/admin/api
RUN go get
EXPOSE 8080
