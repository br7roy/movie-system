#!/bin/bash

masterDep() {

  dir=$(
    cd -P $(dirname $0)
    pwd
  )
  pdir=${dir%/*}
  echo rootDir is $pdir
  rm -fr $pdir/bin

  go build -o $pdir/bin/movie $pdir

  mv -fv $pdir/* ~/.master_deploy
  cd ~/.master_deploy

  git add -A

  ls -l $pdir

  git commit -am "building artifact on $(date +"%Y-%m-%d %H:%M:%S") by Travis-cli"
  git push origin master:master --force --quiet

}

goTest() {

  go test -v $(pwd)

}
run() {
  starttime=$(date +'%Y-%m-%d %H:%M:%S')
  echo -e "start go test...\nstart at:${starttime}"
  goTest
  echo -e "go test done.\nstart deploy master."
  masterDep
  endtime=$(date +'%Y-%m-%d %H:%M:%S')
  start_seconds=$(date --date="$starttime" +%s)
  end_seconds=$(date --date="$endtime" +%s)
  echo -e "all done\nFinish at:${endtime}\nelapse: $((end_seconds - start_seconds))s"
}
run
