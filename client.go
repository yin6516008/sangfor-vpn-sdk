package vpn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

type VpnClient struct {
	host   string
	secret string
	time   string
	client *http.Client
}

func NewVpnClient(host, secret string) *VpnClient {
	return &VpnClient{
		host:   host,
		secret: secret,
		time:   strconv.FormatInt(time.Now().Unix(), 10),
		client: &http.Client{},
	}
}

func ParamsToSortQuery(params map[string]string) string {

	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var resultList []string
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		value := params[key]
		resultList = append(resultList, fmt.Sprintf("%s=%s", key, value))
	}

	result := strings.Join(resultList, "&")
	return result
}

// VpnPost vpn接口post请求基础模版
func (c *VpnClient) VpnPost(uri, reqBody string) ([]byte, error) {
	req, err := http.NewRequest("POST", c.host+uri, strings.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Charset", "UTF-8")
	body, err := c.FormHttp(req)
	return body, err
}
func (c *VpnClient) FormHttp(req *http.Request) ([]byte, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
func NewCommonParams(action, controler string) map[string]string {
	p := map[string]string{
		"action":    action,
		"controler": controler,
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	}
	return p
}

func HandleParams(params interface{}, p map[string]string) (map[string]string, error) {
	paramsByte, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	var mi map[string]interface{}
	err = json.Unmarshal(paramsByte, &mi)
	if err != nil {
		return nil, err
	}

	for k, v := range mi {
		vt := reflect.TypeOf(v)
		switch vt.Kind() {
		case reflect.Map, reflect.Array, reflect.Slice:
			value, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			p[k] = string(value)
		default:
			p[k] = v.(string)
		}
	}
	return p, err
}

func (v *VpnClient) GetSignMap(params interface{}, p map[string]string) map[string]string {
	handleParams, err := HandleParams(params, p)
	if err != nil {
		return nil
	}

	query := ParamsToSortQuery(handleParams)
	var m = map[string]string{
		"sinfor_apitoken": Sha256(query + v.time + v.secret),
		"timestamp":       v.time,
	}
	return m
}
