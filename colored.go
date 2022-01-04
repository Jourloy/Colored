package colored

import "errors"

var (
	FG_Black   = "\x1b[30m"
	FG_Red     = "\x1b[31m"
	FG_Green   = "\x1b[32m"
	FG_Yellow  = "\x1b[33m"
	FG_Blue    = "\x1b[34m"
	FG_Magenta = "\x1b[35m"
	FG_Cyan    = "\x1b[36m"
	FG_White   = "\x1b[37m"
	BG_Black   = "\x1b[40m"
	BG_Red     = "\x1b[41m"
	BG_Green   = "\x1b[42m"
	BG_Yellow  = "\x1b[43m"
	BG_Blue    = "\x1b[44m"
	BG_Magenta = "\x1b[45m"
	BG_Cyan    = "\x1b[46m"
	BG_White   = "\x1b[47m"
	Reset      = "\x1b[0m"
	Bright     = "\x1b[1m"
	Dim        = "\x1b[2m"
	Underscore = "\x1b[4m"
	Blink      = "\x1b[5m"
	Reverse    = "\x1b[7m"
	Hidden     = "\x1b[8m"
	Space      = ""
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
func ChangeFG(t string, FG string) (string, error) {
	if FG == "black" {
		return FG_Black + t + Reset, nil
	} else if FG == "red" {
		return FG_Red + t + Reset, nil
	} else if FG == "green" {
		return FG_Green + t + Reset, nil
	} else if FG == "yellow" {
		return FG_Yellow + t + Reset, nil
	} else if FG == "blue" {
		return FG_Blue + t + Reset, nil
	} else if FG == "magenta" {
		return FG_Magenta + t + Reset, nil
	} else if FG == "cyan" {
		return FG_Cyan + t + Reset, nil
	} else if FG == "white" {
		return FG_White + t + Reset, nil
	} else {
		return "", errors.New("this FG not found")
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
func ChangeBG(t string, BG string) (string, error) {
	if BG == "black" {
		return BG_Black + t + Reset, nil
	} else if BG == "red" {
		return BG_Red + t + Reset, nil
	} else if BG == "green" {
		return BG_Green + t + Reset, nil
	} else if BG == "yellow" {
		return BG_Yellow + t + Reset, nil
	} else if BG == "blue" {
		return BG_Blue + t + Reset, nil
	} else if BG == "magenta" {
		return BG_Magenta + t + Reset, nil
	} else if BG == "cyan" {
		return BG_Cyan + t + Reset, nil
	} else if BG == "white" {
		return BG_White + t + Reset, nil
	} else {
		return "", errors.New("this BG not found")
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
		return Bright + t + Reset, nil
	} else if O == "dim" {
		return Dim + t + Reset, nil
	} else if O == "underscore" {
		return Underscore + t + Reset, nil
	} else if O == "blink" {
		return Blink + t + Reset, nil
	} else if O == "reverse" {
		return Reverse + t + Reset, nil
	} else if O == "hidden" {
		return Hidden + t + Reset, nil
	} else {
		return "", errors.New("this option not found")
	}
}
