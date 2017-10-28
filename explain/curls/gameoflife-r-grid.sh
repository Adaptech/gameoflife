#! /bin/bash
curl -X GET \
  http://127.0.0.1:3001/api/v1/r/grid \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json'
