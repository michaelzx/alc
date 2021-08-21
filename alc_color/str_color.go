package alc_color

func Str(color Color, text string) string {
	return color + text + Reset
}
func RedStr(text string) string {
	return Str(Red, text)
}
func GreenStr(text string) string {
	return Str(Green, text)
}
func YellowStr(text string) string {
	return Str(Yellow, text)
}
func BlueStr(text string) string {
	return Str(Blue, text)
}
func PurpleStr(text string) string {
	return Str(Purple, text)
}
func CyanStr(text string) string {
	return Str(Cyan, text)
}
func WhiteStr(text string) string {
	return Str(White, text)
}
func LightRedStr(text string) string {
	return Str(LightRed, text)
}
func LightGreenStr(text string) string {
	return Str(LightGreen, text)
}
func LightYellowStr(text string) string {
	return Str(LightYellow, text)
}
func LightBlueStr(text string) string {
	return Str(LightBlue, text)
}
func LightPurpleStr(text string) string {
	return Str(LightPurple, text)
}
func LightCyanStr(text string) string {
	return Str(LightCyan, text)
}
func LightWhiteStr(text string) string {
	return Str(LightWhite, text)
}
