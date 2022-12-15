FROM ubuntu/postgres

RUN apt-get -y update

CMD [ "echo", "test works" ]