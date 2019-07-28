package httpreq

import "encoding/json"

// APIFormat API接口
type APIFormat struct {
	Errno  int
	Errmsg error
	Data   interface{}
}

// OkJSON 接口正常的json序列化
func (af *APIFormat) OkJSON() (jsonstr string) {
	af.Errno = 0
	body, err := json.Marshal(af)
	if err != nil {
		return "json encoding error."
	}
	jsonstr = string(body)
	return
}

// ErrJSON 接口异常的json序列化
func (af *APIFormat) ErrJSON() (jsonstr string) {
	if af.Errno == 0 {
		af.Errno = -1
	}
	body, err := json.Marshal(af)
	if err != nil {
		return "json encoding error."
	}
	jsonstr = string(body)
	return
}
