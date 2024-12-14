package utils

import "encoding/json"

// ToJSONString 将任意对象序列化为JSON字符串
func ToJSONString(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToJSONBytes 将任意对象序列化为JSON字节数组
func ToJSONBytes(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// MustToJSONString 将任意对象序列化为JSON字符串,如果出错则panic
func MustToJSONString(v interface{}) string {
	str, err := ToJSONString(v)
	if err != nil {
		panic(err)
	}
	return str
}

// MustToJSONBytes 将任意对象序列化为JSON字节数组,如果出错则panic
func MustToJSONBytes(v interface{}) []byte {
	bytes, err := ToJSONBytes(v)
	if err != nil {
		panic(err)
	}
	return bytes
}

// FromJSONString 将JSON字符串反序列化为任意对象
func FromJSONString(str string, v interface{}) error {
	return json.Unmarshal([]byte(str), v)
}

// FromJSONBytes 将JSON字节数组反序列化为任意对象
func FromJSONBytes(bytes []byte, v interface{}) error {
	return json.Unmarshal(bytes, v)
}

// MustParseJSONString 将JSON字符串反序列化为任意对象,如果出错则panic
func MustParseJSONString[T any](s string) T {
	var result T
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		panic(err)
	}
	return result
}
