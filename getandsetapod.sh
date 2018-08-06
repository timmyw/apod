#!/bin/zsh

echo "START: " $UID $EUID $DESKTOP_SESSION

desktop=mate

case $desktop in
    mate)
        PID=$(pgrep mate-session)
        export DBUS_SESSION_BUS_ADDRESS=$(grep -z DBUS_SESSION_BUS_ADDRESS /proc/$PID/environ|cut -d= -f2-)
    ;;
esac

TMPDIR="$HOME/apod"
mkdir -p $TMPDIR

export PATH=$HOME/go/bin:$PATH
echo $PATH

IMGPATH=`apod --output $TMPDIR`
echo $IMGPATH

        dconf write  /org/mate/desktop/background/picture-filename "'$IMGPATH'"
        dconf write  /org/mate/desktop/background/picture-options "'stretched'"

case $DESKTOP_SESSION in
    mate)
        dconf write  /org/mate/desktop/background/picture-filename "'$IMGPATH'"
        dconf write  /org/mate/desktop/background/picture-options "'stretched'"
        ;;
esac


echo "DONE" $DESKTOP_SESSION
