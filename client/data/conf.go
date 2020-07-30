package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Name         string `json:"name"`
	ServerIpAddr string `json:"server_ip_addr"`
	LocalPort    string `json:"local_port"`
}

func GetConfig() Config {
	conf := Config{}
	//下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	err := Read("./config.json", &conf)
	if err != nil {
		panic(err)
	}
	return conf
}

func Save(con Config) error {
	return Write("./config.json", con)
}

func Read(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func Write(filename string, v interface{}) error {
	body, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, body, os.FileMode(1))
}
