#!/bin/sh
# migrate-db.sh

set -e
  
host="$1"
shift
cmd="$@"

>&2 echo "Begin migrate"
./migrate -path ./schema -database 'postgres://postgres:qwerty@postgres:5432/app-db?sslmode=disable' up
>&2 echo "End migrate"

exec $cmd


