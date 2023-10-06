package gotraceutil

var (
	traceIDKey = DefaultTraceIDKey
	traceKeys  = []string{traceIDKey}
)

const (
	DefaultTraceIDKey = "TraceID"
)

func SetTraceIDKey(key string) {
	traceIDKey = key
	traceKeys[0] = traceIDKey
}

func AppendTraceKeys(keys []string) {
	traceKeys = append(traceKeys, keys...)
}

func ResetTraceKeys() {
	traceIDKey = DefaultTraceIDKey
	traceKeys = []string{traceIDKey}
}

func GetTraceKeys() []string {
	return traceKeys
}
