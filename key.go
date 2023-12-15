package gotraceutil

var (
	tracingIDKey = "TracingID"
	tracingKeys  = []string{tracingIDKey}
)

func GetTracingKeys() []string {
	return tracingKeys
}

func SetTracingKeys(keys []string) {
	if len(keys) == 0 {
		panic("invalid tracingKeys")
	}

	tracingIDKey = keys[0]
	tracingKeys = keys
}
