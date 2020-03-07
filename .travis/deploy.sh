#!/bin/bash
run() {

  dir=$(
    cd -P $(dirname $0)
    pwd
  )
  pdir=${dir%/*}
  echo rootDir is $pdir
  rm -fr $pdir/bin
  starttime=$(date +'%Y-%m-%d %H:%M:%S')
  echo -e "building artifact...\nstart at:${starttime}"
  go build -o $pdir/bin/movie $pdir
  endtime=$(date +'%Y-%m-%d %H:%M:%S')
  start_seconds=$(date --date="$starttime" +%s)
  end_seconds=$(date --date="$endtime" +%s)
  echo -e "building done\nFinish at:${endtime}\nelapse: $((end_seconds - start_seconds))s"

  mv -fv $pdir/* ~/.master_deploy
  cd ~/.master_deploy

  git add -A

  ls -l $pdir

  git commit -am "building artifact on $(date +"%Y-%m-%d %H:%M:%S") by Travis-cli"
  git push origin master:master --force

  echo all done .
}

run
