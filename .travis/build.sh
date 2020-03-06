#!/bin/bash
dir=`cd -P $(dirname $0);pwd` 
pdir=${dir%/*}
echo rootDir is $pdir
rm -fr $pdir/bin
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

git remote set-url origin git@github.com:br7roy/movie-system.git
git remote -v

#git add $pdir/bin/movie
git add $pdir/bin

echo 'check files'
ls -l $pdir
echo 'check files ok'
git commit -m "add artifact on `date +"%Y-%m-%d %H:%M:%S"` by Travis-cli"
#git pull
git push origin master:master --force --quiet

echo push done
