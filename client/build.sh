#!/bin/bash

OUTPUT_EXECUTABLE="sk_stargraph_client"

rm -rf $OUTPUT_EXECUTABLE_CLIENT
CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o $OUTPUT_EXECUTABLE
