/*	данный пакет реализует все потребные для приближения тепловой смерти вселенной функции.
	Методы приближения смерти вселенной можно разделить по источникам:
	– Instagram
	– Vkontakte
	– Telegram
	– Facebook
	– WatsApp
	– Viber
	Для каждого метода есть свой оптимальный размер потому нам и нужно несколько структур
	Собственно соответствующий бот будет принимать интерфейс Heller
	По задумке Рушана Кхаткера параметры размеров картинок надо вынести в файл настроек, а не юать константы
 */
package hell

import (
	"github.com/google/uuid"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)
/* ------------------------------------------------------------- */
const(
	W_Insts	= 1080
	H_Insta = 1080
	W_VkWall = 510
	H_VkWall = 510
	W_Tele = 500
	H_Tele = 500
	Path_hellpng = "../hellpng/"
)

type Heller interface {
	MakeHell() error
	DropHell() error
}
/* ------------------------------------------------------------- */
// внешка для Telegram
type TeleHell struct {
	Hell hell
	Heller
}
func (x *TeleHell) MakeHell() (error) {
	x.Hell.Height = H_Tele
	x.Hell.Width = W_Tele
	return x.Hell.makeHell()
}
func (x *TeleHell) DropHell() error{
	return x.Hell.dropHell()
}
/* ------------------------------------------------------------- */
// внешка для Instagram
type InstaHell struct {
	Hell hell
	Heller
}
func (x *InstaHell) MakeHell() (error) {
	x.Hell.Height = H_Insta
	x.Hell.Width = W_Insts
	return x.Hell.makeHell()
}
func (x *InstaHell) DropHell() error{
	return x.Hell.dropHell()
}
/* ------------------------------------------------------------- */
// внешка для Vkontakte Wall
type VkWallHell struct {
	Hell hell
	Heller
}
func (x *VkWallHell) MakeHell() error {
	x.Hell.Height = H_VkWall
	x.Hell.Width = W_VkWall
	return x.Hell.makeHell()
}
func (x *VkWallHell) DropHell() error{
	return x.Hell.dropHell()
}
/* ------------------------------------------------------------- */
// внутренняя кухня
type hell struct {
	FileName string
	Height int
	Width int
}
func (x *hell) makeHell () error {
	var err error
	// получим уникальную, мать её, строку
	// пусть она станет уникальным, мать его, именем
	x.FileName = Path_hellpng + uuid.New().String() + ".png"
	// Обязательно проверять на существование файла
	// Если файл есть, то не использовать старый, а удалить и создать новый!
	// Больше ада!
	x.dropHell()
	if x.Height * x.Width == 0 {
		// todo обработка ошибки на нулевые размеры
	}
	err = makeHell(x.FileName, x.Width, x.Width)
	return err
}
func (x *hell) dropHell () error {
	if _, err := os.Stat(x.FileName); !os.IsNotExist(err) {
		err = os.Remove(x.FileName)
		if err != nil{
			// todo: файл не удалился по идее надо поместить в лог и что-то делать потом с этой бякой
		}
	}
	return nil
}
/* ------------------------------------------------------------- */
// ну собственно это функция основная, остальное к ней обертка
func makeHell(filename string, width int, height int) (error) {
	// для уникальной, мать её, картинки
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8(rand.Int()),// HUI Warming
				G: uint8(rand.Int()),// Ад в деталях?
				B: uint8(rand.Int()),// АД в рандоме!
				A: 255,
			})
		}
	}
	// да да я скопипастил 99% кода, его спасают от плагиата только избыточные коментарии!
	f, err := os.Create(filename)
	defer f.Close()
	// вот тут надо что-то там проконтролировать и зафаталить
	if err != nil {
		return err
	}
	if err := png.Encode(f, img); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return  err
	}
	return nil
}
