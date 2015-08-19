FROM golang:1.4.2
MAINTAINER tomohiro.teranishi@gmail.com

RUN apt-get update
RUN apt-get install -y vim

RUN go get github.com/tomyhero/battleship-game/...

COPY assets/ /app/assets
RUN chmod 755 /app/assets/entrypoint.sh

ENTRYPOINT ["/app/assets/entrypoint.sh"]
