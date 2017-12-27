package regexRange_test

import (
	"testing"
	"regexp"
	"github.com/ipcjk/regexRange"
)

func TestRangeGenerator(t *testing.T) {
	if regexRange.GetRegex(400, 569) != "_4[0-9][0-9]|_5[0-6][0-9]" {
		t.Error("Failed generating correct regex")
	}
}

func TestPostiveMatch(t *testing.T) {
	regex := regexRange.GetRegex(196608, 197631)
	positiveMatch, err := regexp.Compile(regex)
	if err != nil {
		t.Error("Cant compile regex")
	}

	if !positiveMatch.MatchString("_196922") {
		t.Error("Regex did not match _196922!")
	}
}

func TestNegativeMatch(t *testing.T) {
	/* negative match */
	regex := regexRange.GetRegex(264605, 265628)
	negativeMatch, err := regexp.Compile(regex)
	if err != nil {
		t.Error("Cant compile regex")
	}

	if negativeMatch.MatchString("_196922") {
		t.Error("Regex did match 196922!")
	}
}
