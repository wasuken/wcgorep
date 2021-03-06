package text

import (
	"fmt"
	"github.com/wasuken/wcgorep/common"
	"regexp"
	"strings"
)

type TextMatchResult struct {
	TextMatchLines []TextMatchLine
	Url            string
}

type TextMatchLine struct {
	Number int
	Text   string
}

func (tmr TextMatchResult) Format() {
	fmt.Println("[" + tmr.Url + "]")
	for _, tml := range tmr.TextMatchLines {
		tml.Format()
	}
}

func (tml TextMatchLine) Format() {
	s := fmt.Sprintf("Number: %d, Text: '%s'.", tml.Number, tml.Text)
	fmt.Println(s)
}

func textMatch(contents, pattern string) []TextMatchLine {
	lines := strings.Split(contents, "\n")

	mLines := []TextMatchLine{}

	r := regexp.MustCompile(pattern)
	for i, line := range lines {
		if r.MatchString(line) {
			mLines = append(mLines, TextMatchLine{Number: i, Text: line})
		}
	}

	return mLines
}

func Gorep(url, pattern string) (TextMatchResult, error) {
	c, err := common.Wget(url)
	if err != nil {
		return TextMatchResult{}, err
	}
	return TextMatchResult{Url: url, TextMatchLines: textMatch(c, pattern)}, nil
}
