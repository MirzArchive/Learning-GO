#!/bin/bash


cd exercises/noval_agung || exit

for dir in *; do
  workDir="${dir}"
done

cd "${workDir}" && go run .
