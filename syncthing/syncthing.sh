#!/usr/bin/env bash

type screen >/dev/null 2>&1 || { echo -e >&2 "I require screen but it's not installed. Aborting.\nyou can try:\n\tMac: brew install screen \n\tLinux: \n\t\tapt-get install screen\n\t\tyum install screen"; exit 1; }

ps -efww | grep -w 'syncthing' | grep -v grep | grep -v "$(basename $0)"| awk '{print $2}' | xargs kill -9

echo -e "you will click: [⮐  | return | enter]\nyou can use: [ctrl + a + d] to detach the screen\nyou can run exec: [screen -Ur syncthing] to foreground" ; sleep 5

screen -UR syncthing bash -c "syncthing --no-browser --no-default-folder"