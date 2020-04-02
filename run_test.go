package main

import (
	"fmt"
	"moviex/app/model"
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
