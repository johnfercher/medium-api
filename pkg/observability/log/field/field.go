package field

import "strconv"

type Field struct {
	Key   string
	Value string
}

func Error(value error) *Field {
	return &Field{Key: "error", Value: value.Error()}
}

func StatusCode(value int) *Field {
	return &Field{Key: "status_code", Value: strconv.Itoa(value)}
}

func String(key string, value string) *Field {
	return &Field{Key: key, Value: value}
}

func Int(key string, value int) *Field {
	return &Field{Key: key, Value: strconv.Itoa(value)}
}

func Bool(key string, value bool) *Field {
	return &Field{Key: key, Value: strconv.FormatBool(value)}
}
