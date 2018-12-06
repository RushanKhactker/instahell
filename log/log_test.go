package log

import (
	"fmt"
	"testing"
)

// ну вот нравиться мне адову тьму писать потому пусть падает по приказу!
const DOTESTFAIL = false

// ниже две строки - поебень
func Test_0a(t *testing.T) {
	fmt.Print("============== Это ручной трактора начало \n")
}
func Test_(t *testing.T) {
	if DOTESTFAIL {
		t.Error("Йя уронил этот тест нарочно! не сошел с ума - так нада")
	}
}

// ------------- и да будет тест тут

func TestConsoleLog_Init(t *testing.T) {
	cl := ConsoleLog{}
	cl.Init()
	// проверка на создаваемость
	if &cl.logger == nil {
		t.Error("Ой ... Не смогла я инициализироваться! ")
	}
	// что какой еще может быть проверка ?
}

func TestConsoleLog_Log(t *testing.T) {
	cl := ConsoleLog{}
	cl.Init()
	cl.Log("Какой-то там тестовый лог")
	// а я хз как тут проверить оно верно отработало или нет
}

// ------------- и да завершиться тест тут

// ниже тоже поебень
func Test_0z(t *testing.T) {
	fmt.Print("============== Это ручкой трактора конец \n")
}
