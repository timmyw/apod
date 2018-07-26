#!/bin/sh

dconf write  /org/mate/desktop/background/picture-filename "'$1'"
dconf write  /org/mate/desktop/background/picture-options "'stretched'"
