package common

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const CACHEDIR string = "/tmp/wcgorep/"

func cacheSave(url, contents string) error {
	urlHash := fmt.Sprintf("%x", sha1.Sum([]byte(url)))

	return ioutil.WriteFile(CACHEDIR+urlHash, []byte(contents), 0777)
}

func cacheLoad(url string) (string, bool) {
	if _, err := os.Stat(CACHEDIR); os.IsNotExist(err) {
		if os.Mkdir(CACHEDIR, 0777) != nil {
			panic(err)
		}
	}

	urlHash := fmt.Sprintf("%x", sha1.Sum([]byte(url)))

	if _, err := os.Stat(CACHEDIR + urlHash); os.IsNotExist(err) {
		return "", false
	}
	contents, err := ioutil.ReadFile(CACHEDIR + urlHash)
	if err != nil {
		panic(err)
	}
	return string(contents), true
}

func Wget(url string) (string, error) {
	body, ext := cacheLoad(url)
	if !ext {
		resp, err := http.Get(url)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		body = string(b)
		err = cacheSave(url, body)
		if err != nil {
			return "", err
		}
	}

	return string(body), nil
}
