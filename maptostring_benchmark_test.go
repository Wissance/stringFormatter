package stringFormatter_test

import (
	"testing"

	"github.com/wissance/stringFormatter"
)

func BenchmarkMapToStringWith11Keys(b *testing.B) {
	optionsMap := map[string]any{
		"timeoutMS":                2000,
		"connectTimeoutMS":         20000,
		"maxPoolSize":              64,
		"replicaSet":               "main-set",
		"maxIdleTimeMS":            30000,
		"socketTimeoutMS":          400,
		"serverSelectionTimeoutMS": 2000,
		"heartbeatFrequencyMS":     20,
		"tls":                      "certs/my_cert.crt",
		"w":                        true,
		"directConnection":         false,
	}

	for i := 0; i < b.N; i++ {
		_ = stringFormatter.MapToString(optionsMap, "{key} : {value}", ", ")
	}
}
