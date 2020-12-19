package wcgorep

import (
	"regexp"
	"testing"
)

func TestWget(t *testing.T) {
	url := "https://londone.net"
	result, err := wget(url)
	expected := `
<html>
  <body>
    <p>
      Blog: <a href="https://blog.londone.net">Blog</a>
    </p>
    <p>
      Markdown Space: <a href="https://doc.londone.net">doc</a>
    </p>
    <p>
      Github: <a href="https://github.com/wasuken">Github(wasuken)</a>
    </p>
  </body>
</html>
`
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	r := regexp.MustCompile(`\s+`)

	expected = r.ReplaceAllString(expected, "")
	result = r.ReplaceAllString(result, "")

	if result != expected {
		t.Fatalf("failed test(not equal contents)")
	}
}

func TestWgetFail(t *testing.T) {
	url := "https://londone.net"
	result, err := wget(url)
	expected := "ldasfkadksljffadj"

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	r := regexp.MustCompile(`\s+`)
	expected = r.ReplaceAllString(expected, "")
	result = r.ReplaceAllString(result, "")

	if result == expected {
		t.Fatalf("failed test(equal contents)")
	}
}

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
