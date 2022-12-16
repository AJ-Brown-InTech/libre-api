#default os but development can be in any os
FROM ubuntu:22.04 

#add admin user and dir
RUN apt-get -y update
RUN useradd -m admin && apt-get install -y fish && apt-get install -y golang-go
#create api dir
RUN cd home && cd admin && mkdir api && cd api
# #bundle app source
COPY . ./home/admin/api

#run following cmd in specific api dir
WORKDIR /home/admin/api
#RUN go run main.go
CMD ["echo","complete"]