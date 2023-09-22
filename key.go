package gotraceutil

var (
	traceIDKey string = "TraceId"
	traceKeys         = []string{traceIDKey}
)

func SetTraceIDKey(key string) {
	traceIDKey = key
	traceKeys[0] = traceIDKey
}

func AppendTraceKeys(keys []string) {
	traceKeys = append(traceKeys, keys...)
}

func GetTraceKeys() []string {
	return traceKeys
}
