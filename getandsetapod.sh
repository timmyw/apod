#!/bin/sh

TMPDIR="$HOME/apod"
mkdir -p $TMPDIR

IMGPATH=`apod --output $TMPDIR`

case $DESKTOP_SESSION in
    mate)
        dconf write  /org/mate/desktop/background/picture-filename "'$IMGPATH'"
        dconf write  /org/mate/desktop/background/picture-options "'stretched'"
        ;;
esac


