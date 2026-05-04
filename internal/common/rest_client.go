package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/s4mn0v/listen-trading-api/config"
	"github.com/s4mn0v/listen-trading-api/constants"
)

type BitgetRestClient struct {
	ApiKey     string
	SecretKey  string
	Passphrase string
}

func (p *BitgetRestClient) Init() *BitgetRestClient {
	p.ApiKey = config.GetApiKey()
	p.SecretKey = config.GetSecretKey()
	p.Passphrase = config.GetPassphrase()
	return p
}

func (p *BitgetRestClient) DoGet(uri string, params map[string]string) (string, error) {
	queryString := ""
	if len(params) > 0 {
		queryString = "?"
		for k, v := range params {
			queryString += k + "=" + v + "&"
		}
		queryString = queryString[:len(queryString)-1]
	}

	fullUrl := config.BaseUrl + uri + queryString
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)

	// V2 signature logic
	preHash := timestamp + constants.GET + uri + queryString
	sign := p.generateSignature(preHash)

	req, _ := http.NewRequest(constants.GET, fullUrl, nil)

	req.Header.Add(constants.BgAccessKey, p.ApiKey)
	req.Header.Add(constants.BgAccessSign, sign)
	req.Header.Add(constants.BgAccessPassphrase, p.Passphrase)
	req.Header.Add(constants.BgAccessTimestamp, timestamp)
	req.Header.Add(constants.ContentType, constants.ApplicationJson)
	req.Header.Add("locale", "en-US")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

func (p *BitgetRestClient) generateSignature(message string) string {
	h := hmac.New(sha256.New, []byte(p.SecretKey))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
