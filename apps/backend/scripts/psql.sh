#!/bin/bash

DATABASE=$1

docker exec -it postgresdb psql -U user1 -d $DATABASE