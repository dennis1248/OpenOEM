#!/bin/bash

sh ./buildSetup.sh

# Copy the example config file
cp -f ../config.example.json ../build/config.json

# echo some notes
echo "NOTE: You have just created a release version that is ready to be released on github"
echo "NOTE: If you want to create a new git release create a zip with the config.json and setup.exe from the build dir"