#!/bin/bash

FRAME_LENGTH=4
CAM_DELAY=1

ffmpeg -h > /dev/null 2>&1 || { echo 'ffmpeg not installed, please install it to continue: https://www.ffmpeg.org/' ; exit 1; }
eog -h > /dev/null 2>&1 || { echo 'eye of gnome (eog) not installed' ; exit 1; }
curl -h > /dev/null 2>&1 || { echo 'curl not installed' ; exit 1; }

( timeout $FRAME_LENGTH eog --fullscreen "/home/mmcw/Desktop/ArisaWinter400-24-sample.jpg" &
  (sleep $CAM_DELAY && curl 'https://api.ipify.org?format=json' ) & ) 

#linux eye of gnome app
#timeout 2 eog --fullscreen <file.tiff>

#osx preview app (with applescript arguments?)
#open -a Preview file.tiff