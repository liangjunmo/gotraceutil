package gotraceutil

const (
	DefaultTracingIDKey = "TracingID"
)

var (
	tracingIDKey = DefaultTracingIDKey
	tracingKeys  = []string{tracingIDKey}
)

func GetTracingIDKey() string {
	return tracingIDKey
}

func SetTracingIDKey(key string) {
	tracingIDKey = key
	tracingKeys[0] = tracingIDKey
}

func GetTracingKeys() []string {
	return tracingKeys
}

func AppendTracingKeys(keys []string) {
	tracingKeys = append(tracingKeys, keys...)
}
