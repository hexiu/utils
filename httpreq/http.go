package httpreq

import (
	"net/http"
	"net/url"
)

func Get(urlstr string, params url.Values) (resp *http.Response, err error) {
	urlstr = urlstr + "?" + params.Encode()
	return http.Get(urlstr)
}

func Post(url string, param url.Values) (resp *http.Response, err error) {
	return http.PostForm(url, param)
}

// GetAddHeader 请求添加header头部
func GetAddHeader(urlstr string, params url.Values, header http.Header) (resp *http.Response, err error) {
	urlstr = urlstr + "?" + params.Encode()
	clt := http.DefaultClient
	req, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		return nil, err
	}
	req.Header = header
	return clt.Do(req)
}
