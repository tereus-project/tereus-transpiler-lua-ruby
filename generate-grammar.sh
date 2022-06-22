#!/bin/bash

set -e

java -Xmx500M -cp "./bin/antlr-4.10.1-complete.jar" org.antlr.v4.Tool \
    -Dlanguage=Go \
    -no-listener \
    -visitor \
    -o parser \
    *.g4
