FROM ubuntu/postgres

RUN apt-get -y update

CMD [ "echo", "test works" ]

EXPOSE 5000
EXPOSE 5555
EXPOSE 7070