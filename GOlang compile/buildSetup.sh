#!/bin/bash

rsrc -manifest APP.EXE.manifest -o FILE.syso
GOOS="windows" GOARCH="386" go build -o setup.exe