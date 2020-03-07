#!/bin/bash

git remote set-url origin git@github.com:br7roy/movie-system.git
git remote -v

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
echo building done
echo Finish at:$endtime
echo "elapse: "$((end_seconds-start_seconds))"s"

cd $pdir
git add -A

ls -l $pdir
git branch
git checkout master
git commit -am "building artifact on `date +"%Y-%m-%d %H:%M:%S"` by Travis-cli"
git push origin master:master --force

echo push done
