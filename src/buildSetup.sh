#!/bin/bash
# The build script to create an .exe from the go files 
# Use this package when doing devolopment and just want to test the application

# Create a manifest file for windows
rsrc -manifest APP.EXE.manifest -o FILE.syso

# Create the program
GOOS="windows" GOARCH="386" go build -o setup.exe

# Create a build dir
mkdir ../build -p

# Copy the build files
cp setup.exe ../build/

# Copy the example config file
cp -f ../config.json ../build/config.json

# cleanup
rm -f setup.exe FILE.syso setup.exe~
