package utils

import (
	"fmt"
	"testing"
)



// ну вот нравиться мне адову тьму писать потому пусть падает по приказу!
const DOTESTFAIL = false

// ниже две строки - поебень
func Test_0a(t *testing.T) {fmt.Print("============== Это ручной трактора начало \n")}
func Test_(t *testing.T) {if DOTESTFAIL {t.Error("Йя уронил этот тест нарочно! не сошел с ума - так нада")}}

// ------------- и да будет тест тут

func Test_tst(t *testing.T) {
	var x = MSSQL
	var y = MYSQL
	if x != "mssql" || y != "mysql" {
		t.Error("Пиздец что вы натворили")
	}
}

// ------------- и да завершиться тест тут

// ниже тоже поебень
func Test_0z(t *testing.T) {fmt.Print("============== Это ручкой трактора конец \n")}

}