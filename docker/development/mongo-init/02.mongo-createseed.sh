#!/usr/bin/env bash

echo 'Creating seed data'

mongoimport \
        --host localhost \
        --db ${APP_MONGO_DB} \
        --collection user \
        --type json \
        --file /docker-entrypoint-initdb.d/mongo-seed.json \
        --jsonArray