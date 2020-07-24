package main

import (
	"fmt"
	"moviex/app/model"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"unsafe"
)

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func TestRun(t *testing.T) {

	user := model.User{ID: 1, Email: "haha@.com", Password: "123"}

	var testStruct = &user
	Len := unsafe.Sizeof(*testStruct)
	testBytes := &SliceMock{
		addr: uintptr(unsafe.Pointer(testStruct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(testBytes))
	fmt.Println("[]byte is : ", data)

}
func TestFile(t *testing.T) {
	var fileAath = "C:\\Users\\Tak\\Desktop\\gofile"
	/*	fileInfo, err := ioutil.ReadDir(fileAath)
		if err != nil {
			err.Error()
		}*/

	filepath.Walk(fileAath, func(path string, file os.FileInfo, err error) error {
		if file.IsDir() && !strings.Contains(file.Name(), "@eaDir") {
			fmt.Println(file.Name())
		}
		return nil
	})

}
