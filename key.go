package gotraceutil

var (
	traceIdKey string = "TraceId"
	traceKeys         = []string{traceIdKey}
)

func SetTraceIdKey(key string) {
	traceIdKey = key
	traceKeys[0] = traceIdKey
}

func AppendTraceKeys(keys []string) {
	traceKeys = append(traceKeys, keys...)
}

func GetTraceKeys() []string {
	return traceKeys
}
