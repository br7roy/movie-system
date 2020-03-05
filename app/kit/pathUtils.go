package kit

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// 获取根路径
func RootPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Panicln("发生错误", err.Error())
	}
	i := strings.LastIndex(s, "\\")
	path := s[0 : i+1]
	return path
}
