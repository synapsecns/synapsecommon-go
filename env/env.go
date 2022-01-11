package env

import (
	"os"
	"strconv"
	"time"
)

func String(envKey, fallback string) string {
	if value, ok := os.LookupEnv(envKey); ok {
		return value
	}

	return fallback
}

func Bool(envKey string, fallback bool) bool {
	raw := getValue(envKey)
	if raw == nil {
		return fallback
	}

	parsed := parseValue(*raw, fallback)
	if parsed != nil {
		return parsed.(bool)
	}

	return fallback
}

func Int64(envKey string, fallback int64) int64 {
	raw := getValue(envKey)
	if raw == nil {
		return fallback
	}

	parsed := parseValue(*raw, fallback)
	if parsed != nil {
		return parsed.(int64)
	}

	return fallback
}

func Uint64(envKey string, fallback uint64) uint64 {
	raw := getValue(envKey)
	if raw == nil {
		return fallback
	}

	parsed := parseValue(*raw, fallback)
	if parsed != nil {
		return parsed.(uint64)
	}

	return fallback
}

func Float64(envKey string, fallback float64) float64 {
	raw := getValue(envKey)
	if raw == nil {
		return fallback
	}

	parsed := parseValue(*raw, fallback)
	if parsed != nil {
		return parsed.(float64)
	}

	return fallback
}

func Duration(envKey string, fallback time.Duration) time.Duration {
	raw := getValue(envKey)
	if raw == nil {
		return fallback
	}

	parsed := parseValue(*raw, fallback)
	if parsed != nil {
		return parsed.(time.Duration)
	}

	return fallback
}

func getValue(envKey string) *string {
	val, ok := os.LookupEnv(envKey)
	if ok && val != "" {
		return &val
	}

	return nil
}

func parseValue(value string, fallback interface{}) (res interface{}) {
	if value == "" {
		return nil
	}

	switch fallback.(type) {
	case int64:
		n, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			res = n
		}
	case uint64:
		n, err := strconv.ParseUint(value, 10, 64)
		if err == nil {
			res = n
		}
	case float64:
		n, err := strconv.ParseFloat(value, 64)
		if err == nil {
			res = n
		}
	case bool:
		b, err := strconv.ParseBool(value)
		if err == nil {
			res = b
		}
	case time.Duration:
		d, err := time.ParseDuration(value)
		if err == nil {
			res = d
		}
	}

	return
}
