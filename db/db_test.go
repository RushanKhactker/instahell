package db

import (
	"fmt"
	"testing"
	"time"
)

// ну вот нравиться мне адову тьму писать потому пусть падает по приказу!
const DOTESTFAIL = false

// ниже две строки - поебень
func Test_0a(t *testing.T) {fmt.Print("============== Это ручной трактора начало \n")}
func Test_(t *testing.T) {if DOTESTFAIL {t.Error("Йя уронил этот тест нарочно! не сошел с ума - так нада")}}

// ------------- и да будет тест тут

func Test_z(t *testing.T) {
	l := Log {}
	l.Date = time.Now().String()

}

func TestConnect(t *testing.T) {
	Connect()
}

// ------------- и да завершиться тест тут

// ниже тоже поебень
func Test_0z(t *testing.T) {fmt.Print("============== Это ручкой трактора конец \n")}

