package env

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	ErrWrongFormat = fmt.Errorf("the file is not formatted correctly")
)

// Load loads all the environment variables
// present in the file passed
func Load(file string) error {
	f, err := os.Open(file)

	if err != nil {
		return fmt.Errorf("could not open the file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "=")
		if len(line) < 2 {
			return ErrWrongFormat
		}

		key := strings.TrimSpace(line[0])
		value := strings.TrimSpace(strings.Join(line[1:], ""))

		err = os.Setenv(key, value)
		if err != nil {
			return fmt.Errorf("could not set environment variable: %v", err)
		}
	}

	if scanner.Err() != nil {
		return fmt.Errorf("error while scanning the file: %v", err)
	}

	return nil
}

func String(key, defaultValue string) string {
	v, useDefault := envOrDefault(key)
	if useDefault {
		return defaultValue
	}
	return v
}

func Int(key string, defaultValue int) int {
	v, useDefault := envOrDefault(key)
	return toInt(v, defaultValue, useDefault)
}

func Int32(key string, defaultValue int32) int32 {
	v, useDefault := envOrDefault(key)
	return toInt32(v, defaultValue, useDefault)
}

func Int64(key string, defaultValue int64) int64 {
	v, useDefault := envOrDefault(key)
	return toInt64(v, defaultValue, useDefault)
}

func Float32(key string, defaultValue float32) float32 {
	v, useDefault := envOrDefault(key)
	return toFloat32(v, defaultValue, useDefault)
}

func Float64(key string, defaultValue float64) float64 {
	v, useDefault := envOrDefault(key)
	return toFloat64(v, defaultValue, useDefault)
}

func Bool(key string, defaultValue bool) bool {
	v, useDefault := envOrDefault(key)
	return toBool(v, defaultValue, useDefault)
}

func envOrDefault(key string) (string, bool) {
	env := loadEnv(key)
	if env == nil {
		return "", true
	}
	return *env, false
}
