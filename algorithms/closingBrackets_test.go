package algorithms

import (
	"testing"
)

func TestValidClosingBrackets(t *testing.T) {
	preset1 := []rune{'[', '{', '}', ']'}
	preset2 := []rune{'(', '(', ')', '(', ')', ')'}
	preset3 := []rune{'{', ']'}
	preset4 := []rune{'[', '(', ')', ']', ')', ')', '(', ')'}
	preset5 := []rune{'[', ']', '{', '}', '(', '{', '}', ')'}
	if ValidClosingBrackets(preset1) != true {
		t.Fatal("Preset1 should return true, but false was returned")
	}
	if ValidClosingBrackets(preset2) != true {
		t.Fatal("Preset2 should return true, but false was returned")
	}
	if ValidClosingBrackets(preset3) != false {
		t.Fatal("Preset3 should return false, but true was returned")
	}
	if ValidClosingBrackets(preset4) != false {
		t.Fatal("Preset4 should return false, but true was returned")
	}
	if ValidClosingBrackets(preset5) != true {
		t.Fatal("Preset5 should return true, but false was returned")
	}
}
