#!/bin/bash
mkdir -p "$1"
cp template/main.go "$1/"
touch "$1/test.in"

