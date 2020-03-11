package test

import (
	"fmt"
	"io/ioutil"
	"movie-system/app/kit"
	"path"
	"runtime"
	"testing"
)

func TestReadConf(t *testing.T) {

	if _, fileNameWithPath, _, ok := runtime.Caller(1); ok {
		join := path.Join(fileNameWithPath, "app.yaml")
		d, _ := ioutil.ReadFile(join)
		if d == nil {
			t.FailNow()
		}
	}

	config := kit.AppConfig
	if &config == nil {
		fmt.Println("bad")
	}
}
