package utils

import (
	"encoding/json"
	"strconv"
	"time"
)

func StrToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func Uint64ToStr(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func IsStrBlank(s string) bool { return s == "" }

func TimeToSec(tm time.Time) int64 {
	if tm.IsZero() {
		return 0
	}
	return tm.Unix()
}

func StringPtrToString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func Uint64ToInt64Ptr(i uint64) *int64 {
	v := int64(i)
	return &v
}

func MustJsonDecodeMap(b string) map[string]interface{} {
	vs := make(map[string]interface{})
	_ = json.Unmarshal([]byte(b), &vs)
	return vs
}

func MustJsonDecodeList(b string) []interface{} {
	vs := make([]interface{}, 0)
	_ = json.Unmarshal([]byte(b), &vs)
	return vs
}

func GetExtValue[T float64 | string | bool](ext map[string]interface{}, key string) (T, bool) {
	var d T
	data, exist := ext[key]
	if !exist {
		return d, false
	}
	dataT, ok := data.(T)
	if !ok {
		return d, false
	}
	return dataT, true
}
