package alc_color

type Color = string

const (
	Reset       Color = "\033[0m"
	Red         Color = "\033[31m"
	Green       Color = "\033[32m"
	Yellow      Color = "\033[33m"
	Blue        Color = "\033[34m"
	Purple      Color = "\033[35m"
	Cyan        Color = "\033[36m"
	White       Color = "\033[37m"
	LightRed    Color = "\033[91m"
	LightGreen  Color = "\033[92m"
	LightYellow Color = "\033[93m"
	LightBlue   Color = "\033[94m"
	LightPurple Color = "\033[95m"
	LightCyan   Color = "\033[96m"
	LightWhite  Color = "\033[97m"
	// BlueBold    = "\033[34;1m"
	// MagentaBold = "\033[35;1m"
	// RedBold     = "\033[31;1m"
	// YellowBold  = "\033[33;1m"
)
