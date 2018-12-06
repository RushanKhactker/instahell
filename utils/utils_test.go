package utils

import (
	"fmt"
	"testing"
)

func Test_0a(t *testing.T) {
	fmt.Print("============== Это ручной трактора начало \n")
}

func Test_tst(t *testing.T) {
	var x = MSSQL
	var y = MYSQL
	if x != "mssql" || y != "mysql" {
		t.Error("Пиздец что вы натворили")
	}
}

func Test_0z(t *testing.T) {
	fmt.Print("============== Это ручкой трактора конец \n")
}
