package kit

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type appConfig struct {
	Database database `yaml:"database"`
	Jwt      jwt      `yaml:"jwt"`
	Server   server   `yaml:"server"`
}

type database struct {
	Type       string `yaml:"type"`
	StoreHouse string `yaml:"storeHouse"`
	Port       int    `yaml:"port"`
	Url        string `yaml:"url"`
	UserName   string `yaml:"userName"`
	Password   string `yaml:"password"`
}

type jwt struct {
	Secret  string `yaml:"secret"`
	Expired string `yaml:"expired"`
}

type server struct {
	Mode   string `yaml:"mode"`
	Port   string `yaml:"port"`
	Domain string `yaml:"domain"`
}

var (
	AppConfig   appConfig
	Secret      = "asdasd"
	OneDayHours = 60 * 60 * 24
)

func init() {

	filePath, _ := os.Getwd()
	filePath = path.Join(filePath, "app.yaml")

	data, e := ioutil.ReadFile(filePath)
	if e != nil {
		log.Panicln("读取文件发生错误", e.Error())
	}
	if err := yaml.Unmarshal(data, &AppConfig); err != nil {
		log.Panicln("数据解析错误，无法解析数据", err.Error())
	}
	t := appConfig{}
	//把yaml形式的字符串解析成struct类型
	yaml.Unmarshal(data, &t)
	fmt.Printf("check for config:\n")
	d, _ := yaml.Marshal(&t)
	fmt.Println(string(d))

	fmt.Printf("-----\t config ok .continue\n")
}
