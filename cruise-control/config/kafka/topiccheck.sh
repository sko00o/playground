#!/bin/bash

topics=$(kafka-topics.sh --list --bootstrap-server localhost:9092)

for t in $@; do
    printf "%s\n" "$topics" | grep -Eq "^"$t"$" || exit 1
done
