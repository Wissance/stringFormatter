package stringFormatter

import "testing"

func BenchmarkMapToStringWith11Keys(b *testing.B) {
	optionsMap := map[string]interface{}{
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
		_ = MapToString(&optionsMap, KeyValueWithSemicolonSepFormat, ", ")
	}
}
