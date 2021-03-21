#!/bin/bash
go build &&
./test -writers=3 -arr-size=15 -iter-count=5