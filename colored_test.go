package colored_test

import (
	"github.com/Jourloy/Colored"
	"testing"
)

var testLog = "This is test log"

func TestFG(t *testing.T) {

	colorText, err := colored.ChangeFG(testLog, "red")

	if err != nil {
		t.Error("Error in package: ", err)
	} else if colorText == testLog {
		t.Error("Expected not equal string, got", colorText)
	}
}

func TestBG(t *testing.T) {

	colorText, err := colored.ChangeBG(testLog, "red")

	if err != nil {
		t.Error("Error in package: ", err)
	} else if colorText == testLog {
		t.Error("Expected not equal string, got", colorText)
	}
}

func TestOptions(t *testing.T) {

	colorText, err := colored.AddOptions(testLog, "blink")

	if err != nil {
		t.Error("Error in package: ", err)
	} else if colorText == testLog {
		t.Error("Expected not equal string, got", colorText)
	}
}
