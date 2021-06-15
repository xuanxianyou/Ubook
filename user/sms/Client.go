package sms

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Client struct {
	appId     string
	secretKey string
	signType  string
	version   string
	timestamp string
}

func NewClient() *Client {
	c := &Client{}
	c.SetTimestamp(time.Now().UnixNano() / 1e6)
	c.SetVersion("1.0")
	c.SetSignType("md5")
	return c
}
func (c *Client) SetAppId(appId string) *Client {
	c.appId = appId
	return c
}

func (c *Client) SetSecretKey(secretKey string) *Client {
	c.secretKey = secretKey
	return c
}

func (c *Client) SetVersion(version string) *Client {
	c.version = version
	return c
}

func (c *Client) SetSignType(signType string) *Client {
	c.signType = signType
	return c
}

func (c *Client) SetTimestamp(timestamp int64) *Client {
	c.timestamp = fmt.Sprintf("%d", timestamp)
	return c
}

func (c *Client) CreateSignature(data map[string]string) string {
	keys := make([]string, 0)
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	list := make([]string, 0)
	for _, k := range keys {
		list = append(list, fmt.Sprintf("%s=%s", k, data[k]))
	}
	list = append(list, fmt.Sprintf("%s=%s", "key", c.secretKey))

	str := strings.Join(list, "&")
	str = strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(str))))
	return str
}

func (c *Client) Execute(request *Request) (string, error) {
	post := make(map[string]string)
	post["app_id"] = c.appId
	post["method"] = request.GetMethod()
	post["version"] = c.version
	post["timestamp"] = c.timestamp
	post["sign_type"] = c.signType
	post["biz_content"] = request.GetBizContent()
	post["sign"] = c.CreateSignature(post)
	data := url.Values{}
	for name, value := range post {
		data.Set(name, value)
	}
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://api.shansuma.com/gateway.do", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla 5.0 GO-SMS-SDK")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	buf, err := ioutil.ReadAll(res.Body)
	return string(buf), err
}
