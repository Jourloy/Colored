package colored

var (
	fG_Black   = "\x1b[30m"
	fG_Red     = "\x1b[31m"
	fG_Green   = "\x1b[32m"
	fG_Yellow  = "\x1b[33m"
	fG_Blue    = "\x1b[34m"
	fG_Magenta = "\x1b[35m"
	fG_Cyan    = "\x1b[36m"
	fG_White   = "\x1b[37m"
	bG_Black   = "\x1b[40m"
	bG_Red     = "\x1b[41m"
	bG_Green   = "\x1b[42m"
	bG_Yellow  = "\x1b[43m"
	bG_Blue    = "\x1b[44m"
	bG_Magenta = "\x1b[45m"
	bG_Cyan    = "\x1b[46m"
	bG_White   = "\x1b[47m"
	reset      = "\x1b[0m"
	bright     = "\x1b[1m"
	dim        = "\x1b[2m"
	underscore = "\x1b[4m"
	blink      = "\x1b[5m"
	reverse    = "\x1b[7m"
	hidden     = "\x1b[8m"
)

// Change foreground for text
//
// Second parameter: black | red | green | yellow | blue | magenta | cyan | white
func ChangeFG(t string, fG string) string {
	if fG == "black" {
		return fG_Black + t + reset
	} else if fG == "red" {
		return fG_Red + t + reset
	} else if fG == "green" {
		return fG_Green + t + reset
	} else if fG == "yellow" {
		return fG_Yellow + t + reset
	} else if fG == "blue" {
		return fG_Blue + t + reset
	} else if fG == "magenta" {
		return fG_Magenta + t + reset
	} else if fG == "cyan" {
		return fG_Cyan + t + reset
	} else if fG == "white" {
		return fG_White + t + reset
	} else {
		return t
	}
}

// Change background for text
//
// Second parameter: black | red | green | yellow | blue | magenta | cyan | white
func ChangeBG(t string, bG string) string {
	if bG == "black" {
		return bG_Black + t + reset
	} else if bG == "red" {
		return bG_Red + t + reset
	} else if bG == "green" {
		return bG_Green + t + reset
	} else if bG == "yellow" {
		return bG_Yellow + t + reset
	} else if bG == "blue" {
		return bG_Blue + t + reset
	} else if bG == "magenta" {
		return bG_Magenta + t + reset
	} else if bG == "cyan" {
		return bG_Cyan + t + reset
	} else if bG == "white" {
		return bG_White + t + reset
	} else {
		return t
	}
}

// Add some visual options for text
//
// Second parameter: bright | dim | underscore | blink | reverse | hidden
func AddOptions(t string, O string) string {
	if O == "bright" {
		return bright + t + reset
	} else if O == "dim" {
		return dim + t + reset
	} else if O == "underscore" {
		return underscore + t + reset
	} else if O == "blink" {
		return blink + t + reset
	} else if O == "reverse" {
		return reverse + t + reset
	} else if O == "hidden" {
		return hidden + t + reset
	} else {
		return t
	}
}
