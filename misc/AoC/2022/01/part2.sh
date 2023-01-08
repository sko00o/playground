#!/usr/bin/env bash
cd "$(dirname "$0")"

input="test.txt"
curr=0
list=()
while IFS="" read -r line || [ -n "$line" ]; do
    if [ -n "$line" ]; then
        ((curr += line))
    else
        list+=($curr)
        ((curr = 0))
    fi
done <"$input"

top3=$(
    for i in ${list[*]}; do
        echo $i
    done | sort -n -r | head -3
)

sum=0
for i in ${top3[*]}; do
    ((sum += i))
done

echo $sum