#!/usr/bin/env bash

BASEDIR=$(dirname "$0")

pkgs="
code.olapie.com/mobile
"
cd "$BASEDIR" || exit
GO111MODULE=off GOPROXY=direct GOSUMDB=off gomobile bind -v -ldflags="-s -w" -target=ios -o build/Mobile.xcframework $pkgs
cd - || exit
