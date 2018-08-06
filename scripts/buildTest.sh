#!/bin/bash

# run the build setup
sh ./buildSetup.sh

# copy the testing config to the build dir
cp -f ../examples/config.testing.json ../build/config.json

# copy the test wallpaper to the build dir
cp -f ../examples/testWallpaper.png ../build/testWallpaper.png