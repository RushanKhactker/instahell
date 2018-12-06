package main

import (
	"github.com/instahell/conf"
	"github.com/instahell/log"
)

func main() {
	Log(xconf.DatabaseType)
}

// собственно служебные переменные6 может их будет больеш, кто знает
var xlog log.ConsoleLog
var xconf conf.Config

// инициализация переменных в пакете
func init() {
	xlog.Init()
	xconf = conf.LoadConfig("config.json")
}

// спрячем лог в Лог чтобы меньше писать когда будем вызывать логирование
func Log(a interface{}) {
	xlog.Log(a)
}
