package env

import (
	"os"
	"strconv"
)

func toInt(value string, defaultValue int) int {
	v, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return v
}

func toInt32(value string, defaultValue int32) int32 {
	v, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return defaultValue
	}

	return int32(v)
}

func toInt64(value string, defaultValue int64) int64 {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue
	}

	return v
}

func toFloat32(value string, defaultValue float32) float32 {
	v, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return defaultValue
	}

	return float32(v)
}

func toFloat64(value string, defaultValue float64) float64 {
	v, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return defaultValue
	}

	return float64(v)
}

func toBool(value string, defaultValue bool) bool {
	v, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}

	return v
}

func loadEnv(key string) *string {
	env, ok := os.LookupEnv(key)
	if !ok {
		return nil
	}

	if len(env) == 0 {
		return nil
	}

	return &env
}
