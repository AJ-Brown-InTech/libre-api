#ubuntu 22.04 did not work with github workflows for creating a ci pipeline
# thats the only reason 20.04 is being used
FROM ubuntu:22.04 

#add admin user and dir
RUN apt-get -y update
RUN useradd -m admin && apt-get install -y  && apt-get install fish
# # install golang v1.19 (doesn't matter what sub version)
RUN wget  https://go.dev/dl/go1.19.linux-amd64.tar.gz
RUN tar -xvf go1.19.linux-amd64.tar.gz   
RUN mv go /usr/local 
RUN export GOROOT=/usr/local/go 
RUN export GOPATH=$HOME/Projects/Proj1 
RUN export PATH=$GOPATH/bin:$GOROOT/bin:$PATH 
#create api dir
RUN cd home && cd admin && mkdir api && cd api
# #bundle app source
COPY . ./home/admin/api

#run following cmd in specific api dir
WORKDIR /home/admin/api
#RUN go run main.go
CMD ["echo","complete"]