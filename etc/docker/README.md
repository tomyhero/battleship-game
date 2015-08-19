SETUP
====

BUILD
-----
```
git clone git@github.com:tomyhero/battleship-game.git
cd battleship-game/etc/docker
docker build --tag=tomyhero/battleship .
```

RUN
-----
```
docker run --name battleship  -it -d \
    -p 80:80 \
    -p 8080:8080  \
    -p 9090:9090 \ 
    -e SHIP_HOST=example.com \
    tomyhero/battleship bash
```
