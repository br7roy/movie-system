#!/bin/bash
echo $0 
pdir=`cd -P $(dirname $0);pwd` 
echo $pdir
starttime=`date +'%Y-%m-%d %H:%M:%S'` 
echo building artifact start
echo at:$starttime
go build -o linux/movie ../ 
endtime=`date +'%Y-%m-%d %H:%M:%S'` 
start_seconds=$(date --date="$starttime" +%s)
end_seconds=$(date --date="$endtime" +%s)
echo done
echo Finish at:$endtime
echo "elapse: "$((end_seconds-start_seconds))"s"

