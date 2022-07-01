package result

import "encoding/json"

type Result struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResult() Result {
	return Result{}
}

func (r Result) WithMsg(msg string) Result {
	r.Msg = msg
	return r
}

func (r Result) WithData(data interface{}) Result {
	r.Data = data
	return r
}

func (r Result) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}
