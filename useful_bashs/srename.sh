#!/bin/bash
# rename file or list of files by template

files=`ls`
new_name=""

from_temp=$1
to_temp=$2

if [[ -z $from_temp && -z $to_temp ]]; then
    echo "Please, specify from_template and to_template"
    exit
fi

for file in $files; do
if [[ $file == *$from_temp* ]]; then
    new_name=${file/$from_temp/$to_temp}
    mv $file $new_name
    echo "Renamed from $file to $new_name"
fi
done
