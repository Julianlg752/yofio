#!/bin/bash

export PORT=8000
export DB_USER=yofio_user
export DB_PASSWORD=2021Yofio.5347
export HOSTNAME=localhost
export DB_NAME=yofio

go build
./yofio
