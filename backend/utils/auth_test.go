package utils_test

import (
	"luanti-skin-server/backend/utils"
	"testing"
)

var TestIsValidEmailCases = []struct {
	str    string
	result bool
}{
	{"people[at]gmail.com", false},
	{"people@gmail", true},
	{"people@gmail.com", true},
	{"randomstring", false},
}

func TestIsValidEmail(t *testing.T) {
	for _, c := range TestIsValidEmailCases {
		if utils.IsValidEmail(c.str) != c.result {
			t.Errorf("[%s]: expected %t, got %t", c.str, c.result, !c.result)
		}
	}
}
