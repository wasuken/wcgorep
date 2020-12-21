package common

import (
	"crypto/sha1"
	"fmt"
	"github.com/Songmu/go-httpdate"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const CACHEDIR string = "/tmp/wcgorep/"

func cacheCheck(url string) {
	urlHashPath := CACHEDIR + fmt.Sprintf("%x", sha1.Sum([]byte(url)))
	info, err := os.Stat(urlHashPath)
	if err != nil {
		panic(err)
	}
	cacheModTime := info.ModTime()

	resp, err := http.Head(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	urlLastMod, err := httpdate.Str2Time(strings.TrimSpace(resp.Header.Get("Last-Modified")), time.UTC)
	if err != nil {
		panic(err)
	}

	if cacheModTime.After(urlLastMod) {
		os.Remove(urlHashPath)
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		body := string(b)
		err = cacheSave(url, body)
		if err != nil {
			panic(err)
		}
		if cacheSave(url, body) != nil {
			panic(err)
		}
	}
}

func cacheSave(url, contents string) error {
	urlHash := fmt.Sprintf("%x", sha1.Sum([]byte(url)))

	return ioutil.WriteFile(CACHEDIR+urlHash, []byte(contents), 0777)
}

func cacheLoad(url string) string {
	if _, err := os.Stat(CACHEDIR); os.IsNotExist(err) {
		if os.Mkdir(CACHEDIR, 0777) != nil {
			panic(err)
		}
	}

	urlHash := fmt.Sprintf("%x", sha1.Sum([]byte(url)))

	if _, err := os.Stat(CACHEDIR + urlHash); os.IsNotExist(err) {
		return ""
	}
	contents, err := ioutil.ReadFile(CACHEDIR + urlHash)
	if err != nil {
		panic(err)
	}
	return string(contents)
}

func Wget(url string) (string, error) {
	cacheCheck(url)
	body := cacheLoad(url)

	return string(body), nil
}
