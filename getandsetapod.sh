#!/bin/sh

# Try to work out what desktop session we are running - this isn't
# going to work unless X or Wayland is running

#PID=$(pgrep session) # Hopefully this finds the desktop session

# Trying a more robust way of finding the desktop session
PID=$(pgrep -a session|grep -v worker|grep -v run-script| awk '{print $1}')
desktop=$(grep -z DESKTOP_SESSION  /proc/$PID/environ|cut -d= -f2-)

dbg ()
{
    if [ "$DEBUG" != "" ]; then
        echo $1
    fi
}

dbg "desktop:[$desktop]:"
case $desktop in
    mate)
        PID=$(pgrep mate-session)
        export DBUS_SESSION_BUS_ADDRESS=$(grep -z DBUS_SESSION_BUS_ADDRESS /proc/$PID/environ|cut -d= -f2-)
        ;;
    *)
        echo "Unknown desktop:'$desktop'"
esac

TMPDIR="$HOME/apod"
mkdir -p $TMPDIR

export PATH=$HOME/go/bin:$PATH

IMGPATH=`apod --output $TMPDIR`

case $desktop in
    mate)
        dconf write  /org/mate/desktop/background/picture-filename "'$IMGPATH'"
        dconf write  /org/mate/desktop/background/picture-options "'stretched'"
        ;;
esac

