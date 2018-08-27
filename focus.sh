#!/bin/bash

eog -h > /dev/null 2>&1 || { echo 'eog not installed' ; exit 1; }

eog --fullscreen ./focus.png