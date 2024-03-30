#!/bin/bash

read() {
    printf "%b" "rd$1\x0a\x0a" | nc localhost 3009 | cut -c 3-
}

write () {
    printf "%b" "wd{\"key\":\"$1\",\"value\":\"$2\"}\\x0a\\x0a" | nc localhost 3009 | cut -c 3-
}

if [[ $1 == "read" ]]; then
    read $2
elif [[ $1 == "write" ]]; then
    write $2 $3
else 
    echo "command '$1' not found"
fi