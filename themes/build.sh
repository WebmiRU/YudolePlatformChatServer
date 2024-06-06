#!/bin/bash

for dir in  $(find "$(pwd)" -maxdepth 1 -type d -not -name . -not -name themes)
do
    name=$(basename $dir)
    path=/build/themes/$name
    cd $dir
    npm install
    npm run build
    mkdir -p $path
    cp theme.json $path
    cp -r dist/* $path
done
