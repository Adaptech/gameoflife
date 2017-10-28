#! /bin/bash

curl -X POST \
http://127.0.0.1:3001/api/v1/game/play \
-H 'cache-control: no-cache' \
-H 'content-type: application/json' \
-d '{
  "gameId": "gameId-1"
  , "grid": "grid-1"
  }'
