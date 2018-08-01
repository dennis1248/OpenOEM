#!/bin/bash

# check if all commands are supported
if ! [ -x "$(command -v go)" ]; then
  echo 'Error: Golang not installed read the README.md for more info' >&2
  exit 1
fi

# run the build setup
sh ./buildSetup.sh

# Copy the example config file
cp -f ../config.example.json ../build/config.json

# Create release zip in build folder
go run __makerelease__/main.go

# echo some notes
echo "NOTE: You have just created a release version (build/auto_win_10_conf.zip) that is ready to be released on github"