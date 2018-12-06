package conf

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

// удобная фишка для запуска и проверки
// для каждого модуля надо накатать такую заглушку
func Test_LoadConfig(t *testing.T) {
	config := LoadConfig("test_config.json")

	if config.DatabaseType != "mssql" {
		t.Error("Expected [mssql], got ", config.DatabaseType)
	}
}

// ------------- и да завершиться тест тут

// ниже тоже поебень
func Test_0z(t *testing.T) {
	fmt.Print("============== Это ручкой трактора конец \n")
}
