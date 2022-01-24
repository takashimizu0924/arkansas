package bybit

import (
	"arkansas/config"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Bybit struct {
	apiKey     string
	apiSecret  string
	httpClient *http.Client
	backTest   bool
}

func NewBybit(key, secret string, test bool) *Bybit {
	bybitClient := &Bybit{key, secret, &http.Client{}, test}
	return bybitClient
}

func (b *Bybit) PublicRequest(method, apiURL string, params map[string]interface{}, result interface{}) (resp []byte, err error) {
	var baseURL string
	if b.backTest {
		baseURL = config.BaseURL
	} else {
		baseURL = config.TestBaseURL
	}
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var p []string
	for _, k := range keys {
		p = append(p, fmt.Sprintf("%v=%v", k, params[k]))
	}

	param := strings.Join(p, "&")
	fullURL := baseURL + apiURL
	if param != "" {
		fullURL += "?" + param
	}
	binBody := bytes.NewReader(make([]byte, 0))
	request, err := http.NewRequest(method, fullURL, binBody)
	if err != nil {
		log.Println("action=publicRequest ==> ", err.Error())
		return
	}
	response, err := b.httpClient.Do(request)
	if err != nil {
		log.Println("action=publicRequest ==> ", err.Error())
		return
	}
	defer response.Body.Close()

	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("action=publicRequest ==> ", err.Error())
		return
	}

	err = json.Unmarshal(resp, result)
	return
}

func (b *Bybit) SignedRequest(method, apiURL string, params map[string]interface{}, result interface{}) (resp []byte, err error) {
	var baseURL string
	if b.backTest {
		baseURL = config.BaseURL
	} else {
		baseURL = config.TestBaseURL
	}
	times := time.Now().Unix()
	timestamp := strconv.FormatInt(times, 10)

	params["api_key"] = b.apiKey
	params["timestamp"] = timestamp
	params["recv_window"] = "50000000000000"

	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var p []string
	for _, k := range keys {
		p = append(p, fmt.Sprintf("%v=%v", k, params[k]))
	}
	param := strings.Join(p, "&")
	signature := b.getSign(param)
	param += "&sign=" + signature

	fullURL := baseURL + apiURL + "?" + param

	var binBody = bytes.NewReader(make([]byte, 0))
	request, err := http.NewRequest(method, fullURL, binBody)
	if err != nil {
		log.Println("action=signedRequest ==> ", err.Error())
		return
	}
	response, err := b.httpClient.Do(request)
	if err != nil {
		log.Println("action=signedRequest ==> ", err.Error())
		return
	}
	defer response.Body.Close()

	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("action=signedRequest ==> ", err.Error())
		return
	}
	err = json.Unmarshal(resp, result)

	return
}

func (b *Bybit) getSign(params string) string {
	sig := hmac.New(sha256.New, []byte(b.apiSecret))
	sig.Write([]byte(params))
	signature := hex.EncodeToString(sig.Sum(nil))
	return signature
}
