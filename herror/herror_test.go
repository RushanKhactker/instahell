package herror

import (
	"fmt"
	"testing"
)

// ну вот нравиться мне адову тьму писать потому пусть падает по приказу!
const DOTESTFAIL = false
func Test_0a(t *testing.T) {fmt.Print("============== Это ручной трактора начало \n")}
func Test_(t *testing.T) {if DOTESTFAIL {t.Error("Йя уронил этот тест нарочно! не сошел с ума - так нада")	}}
// ------------- и да будет тест тут

func TestHError_Error(t *testing.T) {
	var herr HError
	herr = HError{"herr",fmt.Errorf("herror")}
	s:= herr.Error()
	valid := fmt.Sprintf("хеггог: х %s, еггог: %v","herr", fmt.Errorf("herror"))
	//fmt.Println(valid)
	if s != valid {
		t.Error("Expected [mssql], got ", s)
	}
}
// ------------- и да завершиться тест тут

// ниже тоже поебень
func Test_0z(t *testing.T) {fmt.Print("============== Это ручкой трактора конец \n")}
