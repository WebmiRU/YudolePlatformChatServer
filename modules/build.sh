#!/bin/bash

for dir in  $(find "$(pwd)" -maxdepth 1 -type d -not -name . -not -name modules)
do
    name=$(basename $dir)
    path=/build/modules/$name
    cd $dir
    GOOS=windows GOARCH=amd64 go build -buildvcs=false -o $name.exe .
    mkdir -p $path
    cp $name.exe $path
    cp module.json $path
    if [ -d resources ]; then
        cp -r resources $path
    fi
done
