#!/bin/bash
dir=`cd -P $(dirname $0);pwd` 
pdir=${dir%/*}
echo rootDir is $pdir
starttime=`date +'%Y-%m-%d %H:%M:%S'` 
echo building artifact start
echo at:$starttime
go build -o $pdir/bin/movie $pdir
endtime=`date +'%Y-%m-%d %H:%M:%S'` 
start_seconds=$(date --date="$starttime" +%s)
end_seconds=$(date --date="$endtime" +%s)
echo done
echo Finish at:$endtime
echo "elapse: "$((end_seconds-start_seconds))"s"

