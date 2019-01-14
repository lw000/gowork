// config_test project main.go
package main

import (
	"log"

	"github.com/Unknwon/goconfig"
	"github.com/syyongx/cconf"

	"sync"
)

func IniConfigTest() {

	var (
		cfg   *goconfig.ConfigFile
		err   error
		value string
	)
	cfg, err = goconfig.LoadConfigFile("./config/config.ini")
	if err != nil {
		log.Printf("读取配置文件失败[config.ini]")
		return
	}

	// 获取冒号为分隔符的键值
	value, err = cfg.GetValue("tail", "log_file")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "log_file", err)
	}
	log.Printf("%s > %s:%s", "super", "log_file", value)

	value, err = cfg.GetValue("tail", "value")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "value", err)
	}
	log.Printf("%s > %s:%s", "super", "value", value)

	value, err = cfg.GetValue("super", "key_super2")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "key_super2", err)
	}
	log.Printf("%s > %s:%s", "super", "key_super2", value)
}

func main() {
	IniConfigTest()

	{
		cfg := cconf.New()
		age := cfg.GetInt("age", 18)
		name := cfg.Get("name", "levi").(string)
		log.Printf("%d", age)
		log.Printf(name)

		cfg.Set("email", "default@default.com")
		email := cfg.GetString("email")
		log.Printf(email)
	}

	{
		pool := &sync.Pool{New: func() interface{} { return "hello, levi" }}

		log.Println(pool.Get().(string))
	}
}
