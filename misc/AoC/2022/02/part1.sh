#!/usr/bin/env bash
cd "$(dirname "$0")"

input="test.txt"

ord() {
    LC_CTYPE=C printf '%d' "'$1"
}

base=$(($(ord X) - $(ord A)))

sum=0
while IFS="" read -r line || [ -n "$line" ]; do
    arr=(${line})
    arg0=$(ord ${arr[0]})
    arg1=$(ord ${arr[1]})
    score=$(($(ord ${arr[1]}) - $(ord X) + 1))
    case $(($arg1 - $arg0 - $base)) in
    1 | -2)
        # echo win
        ((score += 6))
        ;;
    0)
        # echo draw
        ((score += 3))
        ;;
    -1 | 2)
        # echo lost
        ;;
    *)
        echo unknown
        ;;
    esac
    ((sum += score))
done <"$input"

echo $sum
