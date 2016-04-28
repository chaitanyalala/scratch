#!/bin/bash
while [ 1 ]; do
	for f in $PWD/*
	do
		echo $f | nc localhost 2000 > /dev/null 2>&1 &
	done
done
