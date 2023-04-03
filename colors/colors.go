package colors

const (
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorPurple  = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m"
	ColorMagenta = "\033[35m"
)

const (
	TxtCyanBold  = "\033[1;36m"
	TxtWhiteBold = "\033[1;37m"
)

func CoReset() string {
	return string(ColorReset)
}

func CoRed() string {
	return string(ColorRed)
}

func CoGreen() string {
	return string(ColorGreen)
}

func CoYellow() string {
	return string(ColorYellow)
}

func CoBlue() string {
	return string(ColorBlue)
}

func CoPurple() string {
	return string(ColorPurple)
}

func CoCyan() string {
	return string(ColorCyan)
}

func CoWhite() string {
	return string(ColorWhite)
}

// The function prefix's and postfix's the string with reset parameter (\033[0m)
func ResetColor(str string) string {
	newStr := string(ColorReset) + str + string(ColorReset)
	return newStr
}

func RedColor(str string) string {
	newStr := string(ColorRed) + str + string(ColorReset)
	return newStr
}

func GreenColor(str string) string {
	newStr := string(ColorGreen) + str + string(ColorReset)
	return newStr
}

func YellowColor(str string) string {
	newStr := string(ColorYellow) + str + string(ColorReset)
	return newStr
}

func BlueColor(str string) string {
	newStr := string(ColorBlue) + str + string(ColorReset)
	return newStr
}

func PurpleColor(str string) string {
	newStr := string(ColorPurple) + str + string(ColorReset)
	return newStr
}

func CyanColor(str string) string {
	newStr := string(ColorCyan) + str + string(ColorReset)
	return newStr
}

func WhiteColor(str string) string {
	newStr := string(ColorWhite) + str + string(ColorReset)
	return newStr
}
