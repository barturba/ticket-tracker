#!/bin/bash

if [ -f .env ]; then
    source .env
fi

cd sql/schema
goose postgresql $DATABASE_URL up