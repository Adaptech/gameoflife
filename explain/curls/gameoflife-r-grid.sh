#! /bin/bash
curl -X GET \
  http://127.0.0.1:8080/api/v1/r/grid \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json'
