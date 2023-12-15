package gotraceutil

var tracingKeys []string

func GetTracingKeys() []string {
	return tracingKeys
}

func SetTracingKeys(keys []string) {
	tracingKeys = keys
}
