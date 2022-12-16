#ubuntu 22.04 did not work with github workflows for creating a ci pipeline
# thats the only reason 20.04 is being used
FROM ubuntu:22.04 

#setup dir and fs
# RUN apt-get -y update
# RUN useradd -m admin


#add admin user and dir
RUN apt-get -y update
RUN useradd -m admin

# # install golang v1.19 (doesn't matter what sub version)

#create api dir
RUN cd home; cd admin;
RUN mkdir api

#run following cmd in specific api dir
WORKDIR /home/admin/api

#bundle app source
COPY ./ /home/admin/api
#RUN go run main.go
CMD ["echo","complete"]