package common

import (
	"crypto/sha1"
	"fmt"
	"os"
	"regexp"
	"testing"
)

func TestWget(t *testing.T) {
	url := "https://londone.net"
	result, err := Wget(url)
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
	result, err := Wget(url)
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

	urlHash := fmt.Sprintf("%x", sha1.Sum([]byte(url)))

	if _, err := os.Stat(CACHEDIR + urlHash); os.IsNotExist(err) {
		t.Fatalf("failed test(not exists url hash file)")
	}
}
