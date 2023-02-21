#!/usr/bin/env sh

echo ">>> select from rank"
curl ${URL} --data-binary @"${DIR}/001.sql"

echo ">>> select from union rank_a and rank_b"
curl ${URL} --data-binary @"${DIR}/002.sql"

echo ">>> select from rank_simple"
curl ${URL} --data-binary @"${DIR}/003.sql"
