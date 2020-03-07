#!/bin/bash

run() {
  starttime=$(date +'%Y-%m-%d %H:%M:%S')
  echo -e "start go test...\nat:${starttime}"
  go test -v $(pwd)
  echo -e "go test done.\nstart deploy master."
  dir=$(
    cd -P $(dirname $0)
    pwd
  )
  pdir=${dir%/*}
  echo rootDir is $pdir
  rm -fr $pdir/bin

  go build -o $pdir/bin/movie $pdir

  mv -f $pdir/* ~/.master_deploy
  cd ~/.master_deploy

  git add -A

  ls -l $pdir

  git commit -am "deploy artifact on $(date +"%Y-%m-%d %H:%M:%S") by Travis-cli"
  git push origin master:master --force --quiet
  endtime=$(date +'%Y-%m-%d %H:%M:%S')
  start_seconds=$(date --date="$starttime" +%s)
  end_seconds=$(date --date="$endtime" +%s)
  echo -e "all done\nFinish at:${endtime}\nelapse: $((end_seconds - start_seconds))s"
}

run
