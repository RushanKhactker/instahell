package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/instahell/conf"
	"github.com/instahell/db"
	"github.com/instahell/log"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	Log(xconf.DatabaseType)
	fmt.Println(uuid.New())
	db.Connect()
}

// собственно служебные переменные6 может их будет больеш, кто знает
var xlog log.ConsoleLog
var xconf conf.Config

// инициализация переменных в пакете
func init() {
	xlog.Init()
}

// спрячем лог в Лог чтобы меньше писать когда будем вызывать логирование
func Log(a interface{}) {
	xlog.Log(a)
}
