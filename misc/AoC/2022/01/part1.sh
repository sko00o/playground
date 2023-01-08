#!/usr/bin/env bash
cd "$(dirname "$0")"

input="test.txt"
max=0
curr=0
while IFS="" read -r line || [ -n "$line" ]; do
    if [ -n "$line" ]; then
        ((curr += line))
    else
        if ((max < curr)); then
            ((max = curr))
        fi
        ((curr = 0))
    fi
done \
    <"$input"

echo $max
