package gotraceutil

var tracingKeys = []string{"TracingID"}

func GetTracingKeys() []string {
	return tracingKeys
}

func SetTracingKeys(keys []string) {
	tracingKeys = keys
}
