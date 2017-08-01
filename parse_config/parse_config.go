package parse_config

import (
	"log"
	"os"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type DBConfig struct {
	Host 			string 	`yaml:"host"`
	Port 			int		`yaml:"port"`
	User 			string	`yaml:"user"`
	Password 		string	`yaml:"password"`
	Schema 			string	`yaml:"schema"`
	Init_conn_num 	int		`yaml:"init_conn_num"`
	Max_conn_num 	int		`yaml:"max_conn_num"`
}


type TConfig struct {
	Server struct {
		Port 			int 	`yaml:"port"`
		Threads_conn 	int		`yaml:"threads_conn"`
		Workers 		int		`yaml:"workers"`
	}

	Log struct {
		Path 	string		`yaml:"path"`
		Level 	string		`yaml:"level"`
	}

	Db struct {
		Write	DBConfig		`yaml:"write"`
		Read 	[]DBConfig		`yaml:"read,flow"`
	}
}

func Parse_config(filePath string) (TConfig, error) {
	log.Println("filename: ", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Open config file failed, filename: ", filePath)
	}
	defer file.Close()

	configFile, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal("Read config file failed.", string(configFile))
	}

	t := TConfig{}
	err = yaml.Unmarshal(configFile, &t)

	return t, err
}



