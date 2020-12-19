package wcgorep

import (
	"github.com/wasuken/wcgorep/text"
)

func Gorep(args []string) {
	url := args[0]
	ptn := args[1]
	rst, err := text.Gorep(url, ptn)
	if err != nil {
		panic(err)
	}
	rst.Format()
}
