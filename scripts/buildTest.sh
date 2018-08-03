#!/bin/bash

# run the build setup
sh ./buildSetup.sh

# copy the testing config to the build dir
cp -f ../config.testing.json ../build/config.json

# copy the test wallpaper to the build dir
cp -f ../src/testWallpaper.png ../build/testWallpaper.png