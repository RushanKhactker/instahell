package hell

import (
	"os"
	"testing"
)


func TestTeleHell_MakeHell(t *testing.T) {
	var ih TeleHell
	if err := ih.MakeHell(); err != nil {
		t.Errorf("Ошибка: %s", err)
	}
	if _, err := os.Stat(ih.Hell.FileName); os.IsNotExist(err) {
		t.Errorf("File %s not found", ih.Hell.FileName)
	} else {
		os.Remove(ih.Hell.FileName)
	}
}
func TestTeleHell_DropHell(t *testing.T) {
	var ih TeleHell
	ih.MakeHell()
	if err := ih.DropHell(); err != nil {
		t.Errorf("Ошибка: %s", err)
	}
}
func TestInstaHell_MakeHell(t *testing.T) {
	var ih InstaHell
	if err := ih.MakeHell(); err != nil {
		t.Errorf("Ошибка: %s", err)
	}
	if _, err := os.Stat(ih.Hell.FileName); os.IsNotExist(err) {
		t.Errorf("File %s not found", ih.Hell.FileName)
	} else {
		os.Remove(ih.Hell.FileName)
	}
}
func TestInstaHell_DropHell(t *testing.T) {
	var ih InstaHell
	ih.MakeHell()
	if err := ih.DropHell(); err != nil {
		t.Errorf("Ошибка: %s", err)
	}
}
func TestVkWallHell_MakeHell(t *testing.T) {
	var ih VkWallHell
	if err := ih.MakeHell(); err != nil {
		t.Errorf("Ошибка: %s", err)
	}
	if _, err := os.Stat(ih.Hell.FileName); os.IsNotExist(err) {
		t.Errorf("File %s not found", ih.Hell.FileName)
	} else {
		os.Remove(ih.Hell.FileName)
	}
}
func TestVkWallHell_DropHell(t *testing.T) {
	var ih VkWallHell
	ih.MakeHell()
	if err := ih.DropHell(); err != nil {
		t.Errorf("Ошибка: %s", err)
	}
}

