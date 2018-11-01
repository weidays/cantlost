package libs

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	// "projects/FlashPay_go/common/lib/tools"
	// "projects/FlashPay_go/common/shared"

	"strings"
	"time"

	"golang.org/x/net/proxy"
)

type ResponseObject struct {
	Code   int
	Data   string
	Header map[string]string
}

type RequestProxy struct {
	Server   string
	User     string
	Password string
	Port     uint
}

type RequestObject struct {
	URL       string
	header    map[string]string
	Data      string
	Proxy     *RequestProxy
	Response  *ResponseObject
	StartTime time.Time
	EndTime   time.Time
	Silent    bool
}

func (reqObj *RequestObject) Json() string {
	data, _ := json.Marshal(reqObj)
	return string(data)
}

func (reqObj *RequestObject) Header() map[string]string {
	return reqObj.header
}

func (reqObj *RequestObject) SetHeader(key, value string) {
	if reqObj.header == nil {
		reqObj.header = make(map[string]string)
	}
	reqObj.header[key] = value
}

func (reqObj *RequestObject) SetFormMap(data map[string]string) error {
	if data == nil {
		return fmt.Errorf("参数为空")
	}
	forms := make(url.Values)
	for k, v := range data {
		// list = append(list, fmt.Sprintf(k, "=", v))
		forms.Add(k, v)
	}
	return reqObj.SetFormData(forms)
}

func (reqObj *RequestObject) SetFormData(data url.Values) error {
	// list := make([]string, 0, len(data))
	// for k, v := range data {
	// 	list = append(list, fmt.Sprintf(k, "=", v))
	// }
	// reqObj.Data = strings.Join(list, "&")
	if data == nil {
		return fmt.Errorf("参数为空")
	}
	reqObj.Data = data.Encode()
	reqObj.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	return nil
}

func (reqObj *RequestObject) SetQueryData(data url.Values) error {
	// list := make([]string, 0, len(data))
	// for k, v := range data {
	// 	list = append(list, fmt.Sprintf(k, "=", v))
	// }
	// reqObj.Data = strings.Join(list, "&")
	if data == nil {
		return fmt.Errorf("参数为空")
	}
	if reqObj.URL == "" {
		return fmt.Errorf("url为空")
	}
	reqObj.URL = fmt.Sprintf("%s?%s", reqObj.URL, data.Encode())
	// reqObj.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	return nil
}

func (reqObj *RequestObject) SetQueryMap(data map[string]string) error {
	if data == nil {
		return fmt.Errorf("参数为空")
	}
	list := make([]string, 0)
	for k, v := range data {
		list = append(list, k+"="+v)
	}
	reqObj.URL = reqObj.URL + "?" + strings.Join(list, "&")

	// reqObj.Data = data.Encode()
	// reqObj.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	return nil
}

func (reqObj *RequestObject) SetJsonData(data interface{}) error {
	str, err := json.Marshal(data)
	if err != nil {
		return err
	}
	reqObj.Data = string(str)
	reqObj.SetHeader("Content-Type", "application/json")
	return nil
}
func (reqObj *RequestObject) SetXMLData(data interface{}, coding string) error {
	gbkHeader := `<?xml version="1.0" encoding="GBK"?>` + "\n"
	str, err := xml.Marshal(data)
	if err != nil {
		return err
	}
	if coding == "GBK" {
		reqObj.Data = gbkHeader + string(str)
		reqObj.SetHeader("Content-Type", "application/xml;charset=GBK")
	} else {
		reqObj.Data = xml.Header + string(str)
		reqObj.SetHeader("Content-Type", "application/xml")
	}
	return nil
}

func (reqObj *RequestObject) Do() (err error) {
	defer func() {
		if e := recover(); e != nil {
			// fmt.Println("Crash:", e, tools.Stack(1))
			// err = fmt.Errorf(, e, tools.Stack(1))
			fmt.Println(err)
		}
		reqObj.EndTime = time.Now()

		if data, e := json.Marshal(reqObj); e != nil {
			fmt.Println("格式化JSON错误:", e)
		} else {
			if !reqObj.Silent {
				fmt.Println(data)
			}
		}
	}()
	reqObj.StartTime = time.Now()
	httpClient := &http.Client{}
	transport, err := reqObj.setProxy()
	if err != nil {
		return err
	}
	if transport != nil {
		httpClient.Transport = transport
	}
	request := reqObj.createRequest()
	for k, v := range reqObj.header {
		request.Header.Set(k, v)
	}
	reqObj.Response = &ResponseObject{}
	resp, err := httpClient.Do(request)
	reqObj.Response.Header = make(map[string]string)
	if err != nil {
		reqObj.Response.Data = err.Error()
		return fmt.Errorf("请求失败:%s", err)
	}
	reqObj.Response.Code = resp.StatusCode
	for k, v := range resp.Header {
		reqObj.Response.Header[k] = v[0]
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("无法读取请求结果:%s", err)
	}
	reqObj.Response.Data = string(b)
	return nil
}

func (reqObj *RequestObject) setProxy() (*http.Transport, error) {
	if reqObj.Proxy != nil {
		auth := &proxy.Auth{
			User:     reqObj.Proxy.User,
			Password: reqObj.Proxy.Password,
		}
		dialer, err := proxy.SOCKS5("tcp", fmt.Sprintf("%s:%d", reqObj.Proxy.Server, reqObj.Proxy.Port), auth, proxy.Direct)
		if err != nil {
			return nil, fmt.Errorf("无法链接到代理服务器:%s", err.Error())
		}
		httpTransport := &http.Transport{}
		httpTransport.Dial = dialer.Dial
		return httpTransport, nil
	}
	return nil, nil
}

func (reqObj *RequestObject) createRequest() *http.Request {
	var body io.Reader
	method := http.MethodGet
	if reqObj.Data != "" {
		method = http.MethodPost
		body = strings.NewReader(reqObj.Data)
	}
	request, err := http.NewRequest(method, reqObj.URL, body)
	if err != nil {
		panic(fmt.Sprint("无法创建请求:", err))
	}
	return request
}

func (reqObj *RequestObject) ParseForm() map[string]string {
	if reqObj.Response == nil {
		return nil
	}
	values, err := url.ParseQuery(reqObj.Response.Data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	data := make(map[string]string)
	for i := range values {
		if len(values) > 0 {
			data[i] = values[i][0]
		} else {
			data[i] = ""
		}
	}
	// list := strings.Split(, "&")
	// values := make(map[string]string)
	// for i := range list {
	// 	kv := strings.Split(list[i], "=")
	// 	if len(kv) != 2 {
	// 		values[kv[0]] = ""
	// 	} else {
	// 		values[kv[0]] = kv[1]
	// 	}
	// }
	return data
}

func (reqObj *RequestObject) ParseJSON() (map[string]interface{}, error) {
	if reqObj.Response == nil {
		return nil, nil
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(reqObj.Response.Data), &m); err != nil {
		return nil, err
	}
	return m, nil
}
