#!/bin/bash

cat <(cat of.hex) - | socat -t20 -T20 -x - TCP:localhost:6633
