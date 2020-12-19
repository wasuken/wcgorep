package text

import (
	"testing"
)

func TestTextMatch(t *testing.T) {
	contents := `hoge
unchi
buriburi
ldajflnfunchikljfalds
dskaf
kljfa
ksd`
	expected := []TextMatchLine{
		TextMatchLine{
			Text:   "unchi",
			Number: 1,
		},
		TextMatchLine{
			Text:   "ldajflnfunchikljfalds",
			Number: 3,
		},
	}
	result := textMatch(contents, "unchi")
	if len(expected) != len(result) {
		t.Fatalf("failed test(not equal matches line)")
	}
	for i, tml := range result {
		if tml.Text != expected[i].Text || tml.Number != expected[i].Number {
			t.Fatalf("failed test(not equal line contents or number)")
		}
	}
}
