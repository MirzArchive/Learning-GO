#!/bin/bash

dirs=$(cd exercises && cd noval_agung && ls --format=horizontal)
IFS=" " read -r -a dirArr <<< "$dirs"

arrLength=${#dirArr[*]}
lastIndex=$(( arrLength - 1 ))

cd "exercises/noval_agung/${dirArr[${lastIndex}]}" && go run .
