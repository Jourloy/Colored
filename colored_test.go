package colored_test

import (
	"github.com/Jourloy/Colored"
	"testing"
)

var testLog = "This is test log"

func TestFG(t *testing.T) {
	if colorText := colored.ChangeFG(testLog, "red"); colorText == testLog {
		t.Error("Expected not equal string, got", colorText)
	}
}

func TestFG_Fail(t *testing.T) {
	if colorText := colored.ChangeFG(testLog, "what"); colorText != testLog {
		t.Error("Expected equal string, got", colorText)
	}
}

func TestBG(t *testing.T) {
	if colorText := colored.ChangeBG(testLog, "red"); colorText == testLog {
		t.Error("Expected not equal string, got", colorText)
	}
}

func TestBG_Fail(t *testing.T) {
	if colorText := colored.ChangeBG(testLog, "what"); colorText != testLog {
		t.Error("Expected equal string, got", colorText)
	}
}

func TestOptions(t *testing.T) {
	if colorText := colored.AddOptions(testLog, "blink"); colorText == testLog {
		t.Error("Expected not equal string, got", colorText)
	}
}

func TestOptions_Fail(t *testing.T) {
	if colorText := colored.ChangeBG(testLog, "what"); colorText != testLog {
		t.Error("Expected equal string, got", colorText)
	}
}
