package alc_print

import "fmt"

var (
	greenBg   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	whiteBg   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellowBg  = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	redBg     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blueBg    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magentaBg = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyanBg    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	green     = string([]byte{27, 91, 51, 50, 109})
	white     = string([]byte{27, 91, 51, 55, 109})
	yellow    = string([]byte{27, 91, 51, 51, 109})
	red       = string([]byte{27, 91, 51, 49, 109})
	blue      = string([]byte{27, 91, 51, 52, 109})
	magenta   = string([]byte{27, 91, 51, 53, 109})
	cyan      = string([]byte{27, 91, 51, 54, 109})
	rest      = string([]byte{27, 91, 48, 109})
)

func GreenBg(str string) {
	fmt.Println(greenBg, str, rest)
}

func WhiteBg(str string) {
	fmt.Println(whiteBg, str, rest)
}
func YellowBg(str string) {
	fmt.Println(yellowBg, str, rest)
}
func RedBg(str string) {
	fmt.Println(redBg, str, rest)
}
func BlueBg(str string) {
	fmt.Println(blueBg, str, rest)
}
func MagentaBg(str string) {
	fmt.Println(magentaBg, str, rest)
}
func CyanBg(str string) {
	fmt.Println(cyanBg, str, rest)
}
func Green(str string) {
	fmt.Println(green, str, rest)
}
func White(str string) {
	fmt.Println(white, str, rest)
}
func Yellow(str string) {
	fmt.Println(yellow, str, rest)
}
func Red(str string) {
	fmt.Println(red, str, rest)
}
func Blue(str string) {
	fmt.Println(blue, str, rest)
}
func Magenta(str string) {
	fmt.Println(magenta, str, rest)
}
func Cyan(str string) {
	fmt.Println(cyan, str, rest)
}
