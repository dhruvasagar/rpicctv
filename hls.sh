#!/bin/bash

raspivid -n -w 720 -h 405 -fps 25 -vf -t 86400000 -b 1800000 -ih -o - \
| ffmpeg -y \
    -i - \
    -c:v copy \
    -map 0:0 \
    -hls_wrap 100 -hls_time 10 -hls_flags delete_segments -hls_list_size 10 -f hls live/live.m3u8
