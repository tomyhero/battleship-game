#!/bin/bash
set -e

SHIP_HOST=${SHIP_HOST:-localhost}
SHIP_MATCH_PORT=${SHIP_MATCH_PORT:-8080}
SHIP_GAME_PORT=${SHIP_GAME_PORT:-9090}
SHIP_HTML5_CLIENT_PORT=${SHIP_HTML5_CLIENT_PORT:-80}

sed s/{{SHIP_HOST}}/$SHIP_HOST/g -i /app/assets/config.toml
sed s/{{SHIP_MATCH_PORT}}/$SHIP_MATCH_PORT/g -i /app/assets/config.toml
sed s/{{SHIP_GAME_PORT}}/$SHIP_GAME_PORT/g -i /app/assets/config.toml

cd /go/src/github.com/tomyhero/battleship-game/ ; git pull ; git checkout develop ; go get github.com/tomyhero/battleship-game/... ; go install ./...

# とりあえず適当に立ち上げる
/go/bin/matching --config="/app/assets/config.toml" --port=${SHIP_MATCH_PORT} &
/go/bin/game --config="/app/assets/config.toml" --port=${SHIP_GAME_PORT} &
/go/bin/html5_client --config="/app/assets/config.toml" --port=${SHIP_HTML5_CLIENT_PORT} &



bash
