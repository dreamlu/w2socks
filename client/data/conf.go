package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dreamlu/w2socks/client/core"
	"io/ioutil"
	"os"
)

type Config struct {
	Name string `json:"name"`
	core.W2Config
}

var Path = "./.w2socks.json"

func GetConfig() []*Config {
	var conf []*Config
	//下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	err := Read(Path, &conf)
	if err != nil {
		return conf
	}
	return conf
}

// 存储配置文件
func InsertConfig(con Config) error {
	conf := GetConfig()
	for _, v := range conf {
		if con.W2Config.String() == v.W2Config.String() {
			return errors.New("the same server_ip_addr and local_port existed")
		}
	}
	conf = append(conf, &con)
	return Write(Path, conf)
}

// 存储配置文件
func UpdateConfig(con Config, index int) error {
	conf := GetConfig()
	for k, v := range conf {
		if con.W2Config.String() == v.W2Config.String() && k != index {
			return errors.New("the same server_ip_addr and local_port existed")
		}
	}
	conf[index] = &con
	return Write(Path, conf)
}

// 存储配置文件
func DeleteConfig(con Config) error {
	conf := GetConfig()
	for k, v := range conf {
		if con.W2Config.String() == v.W2Config.String() {
			conf = append(conf[:k], conf[k+1:]...)
		}
	}
	return Write(Path, conf)
}

func Read(filename string, v interface{}) error {
	Check(filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if string(data) == "" {
		return fmt.Errorf("配置文件为空")
	}
	return json.Unmarshal(data, v)
}

func Write(filename string, v interface{}) error {
	Check(filename)
	body, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, body, 0666)
}

func Check(filename string) {
	exist, err := PathExists(filename)
	if err != nil {
		fmt.Printf("[error!][%v]\n", err)
		return
	}

	// 文件不存在
	if !exist {
		file, err := os.Open(filename)
		defer func() { file.Close() }()
		if err != nil && os.IsNotExist(err) {
			file, err = os.Create(filename)
			if err != nil {
				fmt.Printf("[error!][%v]\n", err)
			}
			file.Close()
		}
	}
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
