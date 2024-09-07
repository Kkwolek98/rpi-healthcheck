#!/bin/bash
if [ "$1" == "measure_temp" ]; then
    rand = awk -v min=45 -v max=60 'BEGIN{srand(); printf "%.2f\n", min+rand()*(max-min)}'
    echo "temp=$rand'C"
    exit 0
fi

echo "Error: unsupported command: $1"
exit 1
