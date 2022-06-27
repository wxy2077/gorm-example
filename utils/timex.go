package utils

import (
	"github.com/guregu/null"
)

type lt = null.Time
type LocalTimeX struct {
	lt
}

// 重写 null.Time 的序列化方法
func (t LocalTimeX) MarshalJSON() ([]byte, error) {

	if !t.Valid {
		return []byte("null"), nil
	}
	tune := t.Time.Format(`"2006-01-02 15:04:05"`)

	return []byte(tune), nil
}
