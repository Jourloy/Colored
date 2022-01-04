package colored

import "errors"

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
// Second parameter:
//
// - black
//
// - red
//
// - green
//
// - yellow
//
// - blue
//
// - magenta
//
// - cyan
//
// - white
func ChangefG(t string, fG string) (string, error) {
	if fG == "black" {
		return fG_Black + t + reset, nil
	} else if fG == "red" {
		return fG_Red + t + reset, nil
	} else if fG == "green" {
		return fG_Green + t + reset, nil
	} else if fG == "yellow" {
		return fG_Yellow + t + reset, nil
	} else if fG == "blue" {
		return fG_Blue + t + reset, nil
	} else if fG == "magenta" {
		return fG_Magenta + t + reset, nil
	} else if fG == "cyan" {
		return fG_Cyan + t + reset, nil
	} else if fG == "white" {
		return fG_White + t + reset, nil
	} else {
		return "", errors.New("this fG not found")
	}
}

// Change background for text
//
// Second parameter:
//
// - black
//
// - red
//
// - green
//
// - yellow
//
// - blue
//
// - magenta
//
// - cyan
//
// - white
func ChangebG(t string, bG string) (string, error) {
	if bG == "black" {
		return bG_Black + t + reset, nil
	} else if bG == "red" {
		return bG_Red + t + reset, nil
	} else if bG == "green" {
		return bG_Green + t + reset, nil
	} else if bG == "yellow" {
		return bG_Yellow + t + reset, nil
	} else if bG == "blue" {
		return bG_Blue + t + reset, nil
	} else if bG == "magenta" {
		return bG_Magenta + t + reset, nil
	} else if bG == "cyan" {
		return bG_Cyan + t + reset, nil
	} else if bG == "white" {
		return bG_White + t + reset, nil
	} else {
		return "", errors.New("this bG not found")
	}
}

// Add some visual options for text
//
// Second parameter:
//
// - bright
//
// - dim
//
// - underscore
//
// - blink
//
// - reverse
//
// - hidden
func AddOptions(t string, O string) (string, error) {
	if O == "bright" {
		return bright + t + reset, nil
	} else if O == "dim" {
		return dim + t + reset, nil
	} else if O == "underscore" {
		return underscore + t + reset, nil
	} else if O == "blink" {
		return blink + t + reset, nil
	} else if O == "reverse" {
		return reverse + t + reset, nil
	} else if O == "hidden" {
		return hidden + t + reset, nil
	} else {
		return "", errors.New("this option not found")
	}
}
