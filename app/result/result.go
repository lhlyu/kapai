package result

import "encoding/json"

type Result[T any] struct {
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func NewResult[T any]() Result[T] {
	return Result[T]{}
}

func (r Result[T]) WithMsg(msg string) Result[T] {
	r.Msg = msg
	return r
}

func (r Result[T]) WithData(data T) Result[T] {
	r.Data = data
	return r
}

func (r Result[T]) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}
