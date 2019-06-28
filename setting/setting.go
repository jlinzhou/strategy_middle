package setting

import (
	//"fmt"
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	//定义*ini.File类型的Cfg
	Cfg *ini.File
	//运行模式
	RunMode string
	//端口
	HTTPPort int
	//time.Duration毫秒
	//读时间
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	//mongodb数据库账户密码 ip及端口
	Mongodb_USER     string
	Mongodb_PASSWORD string
	Mongodb_HOST     string
	Mongodb_DATABASE string

	//rabbitmq账户密码
	Rabbitmq_USER string
	Rabbitmq_PASSWORD string
	Rabbitmq_HOST string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadDatabase()
	LoadRabbitmq()
	// LoadApp()

}

//加载运行模式
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	//取setting值发生错误就取MustString
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(3000)

	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadDatabase() {
	sec, err := Cfg.GetSection("mongodb")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	Mongodb_USER = sec.Key("USER").MustString("jlzhou")
	Mongodb_PASSWORD = sec.Key("PASSWORD").MustString("jlzhou")
	Mongodb_HOST = sec.Key("HOST").MustString("192.168.18.16:27017")
	Mongodb_DATABASE = sec.Key("DATABASE").MustString("test")

}
func LoadRabbitmq() {
	sec, err := Cfg.GetSection("rabbitmq")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	Rabbitmq_USER = sec.Key("USER").MustString("dev")
	Rabbitmq_PASSWORD = sec.Key("PASSWORD").MustString("dev")
	Rabbitmq_HOST = sec.Key("HOST").MustString("192.168.18.24:5672")


}
// func LoadApp() {
// 	sec, err := Cfg.GetSection("app")
// 	if err != nil {
// 		log.Fatalf("Fail to get section 'app': %v", err)
// 	}
// 	//密钥？？
// 	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
// 	//最大页数
// 	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
// }
