package config

import (
	"time"

	"github.com/tkanos/gonfig"
)

type AppConfig struct {
	ServiceName string `json:"serviceName"`
	Key         string `json:"key"`
	Value       string `json:"value"`
}

type DbConfig struct {
	ServiceName string `json:"serviceName"`
	DbHost      string `json:"dbHost"`
	DbPort      int32  `json:"dbPort"`
	DbName      string `json:"dbName"`
	DbUser      string `json:"dbUser"`
	DbPassword  string `json:"dbPassword"`
	DbTimeOut   int32  `json:"dbTimeOut"`
	DbReconnect int32  `json:"dbReconnect"`
}

type ServiceAddrConfig struct {
	ServiceName string `json:"serviceName"`
	Host        string `json:"host"`
	Port        int32  `json:"port"`
}


//Configuration struct: store all common config
type Configuration struct {
	NodeDbServer    string        `json:"nodeDbServer"`
	NodeDbPort      int32         `json:"nodeDbPort"`
	NodeDbName      string        `json:"nodeDbName"`
	NodeDbTimeOut   int32         `json:"nodeDbTimeOut"`
	NodeDbReconnect int32         `json:"nodeDbReconnect"`
	NodeDbUser      string        `json:"nodeDbUser"`
	NodeDbPassword  string        `json:"nodeDbPassword"`
	Debug           bool          `json:"debug"`
	ServerPort      int32         `json:"serverPort"`
	Authenticate    bool          `json:"authenticate"`
	DBServer        string        `json:"dbServer"`
	DBPort          int32         `json:"dbPort"`
	DBName          string        `json:"dbName"`
	AppName         string        `json:"appName"`
	DBTimeOut       int32         `json:"dbTimeOut"`
	DBReconnect     int32         `json:"dbReconnect"`
	DBUser          string        `json:"dbUser"`
	DBPassword      string        `json:"dbPassword"`
	PrivateKey      string        `json:"privateKey"`
	JwtExpDuration  time.Duration `json:"jwtExpDuration"`
	CallTimeout     int32         `json:"callTimeout"`
	ImageRatio      float64       `json:"imageRatio"`
	HtmlToPdfApp    string        `json: "htmlToPdfApp"`
	ChunkSize       int64         `json: "chunkSize"`
	ReportServer    string        `json:"ReportServer"`
	ImageServerPort int32         `json:"imageServerPort"`
	ImageLocation   string        `json:"ImageLocation"`
	PackagePrefix   string        `json:"packagePrefix"`
}

//ServiceAddr struct
type ServiceAddr struct {
	CoreService string `json:"coreService"`
	File        string `json:"file"`
	Report      string `json:"report"`
	Skyins      string `json:"skyins"`
	Skycmn      string `json:"skycmn"`
	Skyinv      string `json:"skyinv"`
	Skyatc      string `json:"skyatc"`
	Skyreg      string `json:"skyreg"`
	Skyimg      string `json:"skyimg"`
	Skyemr      string `json:"skyemr"`
	Skylab      string `json:"skylab"`
	Skyacc      string `json:"skyacc"`
	Skysle      string `json:"skysle"`
	Skyrpt      string `json:"skyrpt"`
	Skylis      string `json:"skylis"`
	Skyris      string `json:"skyris"`
	Skypacs     string `json:"skypacs"`
	Skysur      string `json:"skysur"`
}

//GlobalConfig store configuration globally
var GlobalConfig Configuration
var GlobalServiceAddr ServiceAddr

//LoadConfig load config from config.json file
func LoadConfig(parentPath, configFile string) (conf Configuration, err error) {
	err = gonfig.GetConf(parentPath+"config/"+configFile, &conf)
	conf.JwtExpDuration = conf.JwtExpDuration * time.Second
	GlobalConfig = conf
	return conf, err
}

//LoadServiceAddr
func LoadServiceAddr(parentPath, configFile string) (conf ServiceAddr, err error) {
	err = gonfig.GetConf(parentPath+"config/"+configFile, &conf)
	GlobalServiceAddr = conf
	return conf, err
}

