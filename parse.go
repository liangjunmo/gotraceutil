package gotraceutil

import "context"

func Parse(ctx context.Context) (labels map[string]string) {
	if ctx == nil {
		return nil
	}

	labels = make(map[string]string)

	for _, key := range tracingKeys {
		val := ctx.Value(key)

		if val == nil {
			labels[key] = ""
		} else {
			labels[key] = val.(string)
		}
	}

	return labels
}
