FROM alpine:latest

ADD . /home/app/
WORKDIR /home/app

RUN chmod 777 /home/app/Hitokoto

ENTRYPOINT ["./Hitokoto"]
EXPOSE 8080