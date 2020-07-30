package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Name         string `json:"name"`
	ServerIpAddr string `json:"server_ip_addr"`
	LocalPort    string `json:"local_port"`
}

var Path = "./config.json"

func GetConfig() []*Config {
	var conf []*Config
	//下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	err := Read(Path, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}

// 存储配置文件
func SaveConfig(con Config) error {
	return Write(Path, con)
}

func Read(filename string, v interface{}) error {
	Check(filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	dataStr := string(data)
	if dataStr == "" {
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
