package location

import (
	"github.com/SUCHMOKUO/falcon-ws/util"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var apiURL string

// SetApiUrl set the serverURL of
// the location query function.
func SetApiUrl(url string) {
	apiURL = url
}

// IsInChina detect if the host location is in china
// through the server api.
func IsInChina(host string) bool {
	if apiURL == "" {
		log.Fatalln("should call SetApiUrl first")
	}

	hostEncoded := util.Encode(host)
	url := apiURL + "?h=" + hostEncoded

	resp, err := http.Get(url)
	if err != nil {
		log.Println("query location error:", err)
		return true
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read body of query location error:", err)
		return true
	}

	if isInMainland(string(body)) {
		return true
	}
	return false
}

var regx = regexp.MustCompile(`(台湾|香港|澳门)`)

func isInMainland(location string) bool {
	if !strings.Contains(location, "中国") {
		return false
	}
	return !regx.MatchString(location)
}
