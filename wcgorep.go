package wcgorep

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func wget(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func Gorep(args []string) {
	url := args[0]
	fmt.Println(url)
}
