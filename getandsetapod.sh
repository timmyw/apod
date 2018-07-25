#!/bin/sh

TMPDIR="/tmp"

IMGPATH=`apod --output $TMPDIR`

dconf write  /org/mate/desktop/background/picture-filename "'$IMGPATH'"


