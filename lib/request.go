package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type RpJson struct {
	Url         string
	Method      string
	ContentType string
	Opt         string
	SerViceName string
	SpanName    string
	Spanctx     model.SpanContext
	Kind        model.Kind
}

func (rp *RpJson) RpZipkin() (string, error) {
	var body io.Reader
	if rp.Method == "GET" {
		body = nil
	} else {
		switch rp.ContentType {
		case "text/xml":
			body = strings.NewReader(rp.Opt)
		case "application/x-www-form-urlencoded":
			body = strings.NewReader(rp.Opt)
		case "application/json":
			jsonStr := []byte(rp.Opt)
			body = bytes.NewBuffer(jsonStr)
		default:
		}
	}

	childSpan, _ := GetSpan(rp.SerViceName, rp.SpanName, rp.Spanctx, rp.Kind)
	//defer reporter.Close()
	defer childSpan.Finish()
	param := rp.Opt
	if rp.SpanName != "908008" {
		length := len([]rune(rp.Opt))
		param = string([]rune(rp.Opt))
		if length > 500 {
			param = string([]rune(rp.Opt)[:500])
		}
	}

	childSpan.Tag("request param", param)

	req, err := http.NewRequest(rp.Method, rp.Url, body)

	if err != nil {
		zipkin.TagError.Set(childSpan, err.Error())
		return "", err
	}
	zipkin.TagHTTPMethod.Set(childSpan, req.Method)
	//zipkin.TagHTTPPath.Set(childSpan, req.URL.Path)
	zipkin.TagHTTPUrl.Set(childSpan, req.URL.Host+req.URL.Path)

	if rp.Method != "GET" {
		req.Header.Set("Content-Type", rp.ContentType)
	}
	client := &http.Client{}

	rsp, err := client.Do(req.WithContext(zipkin.NewContext(req.Context(), childSpan)))
	if err != nil {
		zipkin.TagError.Set(childSpan, err.Error())
		return "", err
	}

	if rsp != nil {
		defer rsp.Body.Close()
	}

	data, _ := ioutil.ReadAll(rsp.Body)
	length := len([]rune(string(data)))
	param = string([]rune(string(data)))
	if length > 500 {
		param = string([]rune(string(data))[:500])
	}
	zipkin.TagHTTPStatusCode.Set(childSpan, strconv.Itoa(rsp.StatusCode))
	childSpan.Tag("response", param)

	if rp.SerViceName == "portal_auth" {
		var result map[string]interface{}
		json.Unmarshal([]byte(string(data)), &result)
		if v, ok := result["access_token"]; ok && v != nil {
			token := (result["access_token"]).(map[string]interface{})
			childSpan.Tag("portal_id", strconv.FormatFloat(token["portal_id"].(float64), 'f', -1, 64))
		}
	}

	if rsp.StatusCode == 200 {
		return string(data), nil
	} else {
		return "", errors.New(rsp.Status)
	}

}

func (rp *RpJson) RP() (string, error) {
	result, err := rp.RpZipkin()
	return result, err
}
