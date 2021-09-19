package env

import (
	"log"
	"os"
	"strconv"
)

type EnvParser interface {
	String(key, defaultValue string) string
	Int(key string, defaultValue int) int
	Int32(key string, defaultValue int32) int32
	Int64(key string, defaultValue int64) int64
	Float32(key string, defaultValue float32) float32
	Float64(key string, defaultValue float64) float64
	Bool(key string, defaultValue bool) bool
}

type Parser struct {
	l *log.Logger
}

func NewParser(logger *log.Logger) *Parser {
	return &Parser{logger}
}

func (e *Parser) String(key, defaultValue string) string {
	env := e.loadEnv(key)
	if env == nil {
		return defaultValue
	}

	return *env
}

func (e *Parser) Int(key string, defaultValue int) int {
	env := e.loadEnv(key)
	if env == nil {
		return defaultValue
	}

	v, err := strconv.Atoi(*env)
	if err != nil {
		e.l.Printf("[INFO] %s is set but is not an int, using default value\n", key)
		return defaultValue
	}

	return v
}

func (e *Parser) Int32(key string, defaultValue int32) int32 {
	env := e.loadEnv(key)
	if env == nil {
		return defaultValue
	}

	v, err := strconv.ParseInt(*env, 10, 32)
	if err != nil {
		e.l.Printf("[INFO] %s is set but is not an int32, using default value\n", key)
		return defaultValue
	}

	return int32(v)
}

func (e *Parser) Int64(key string, defaultValue int64) int64 {
	env := e.loadEnv(key)
	if env == nil {
		return defaultValue
	}

	v, err := strconv.ParseInt(*env, 10, 64)
	if err != nil {
		e.l.Printf("[INFO] %s is set but is not an int64, using default value\n", key)
		return defaultValue
	}

	return v
}

func (e *Parser) Float32(key string, defaultValue float32) float32 {
	env := e.loadEnv(key)
	if env == nil {
		return defaultValue
	}

	v, err := strconv.ParseFloat(*env, 32)
	if err != nil {
		e.l.Printf("[INFO] %s is set but is not a float32, using default value\n", key)
		return defaultValue
	}

	return float32(v)
}

func (e *Parser) Float64(key string, defaultValue float64) float64 {
	env := e.loadEnv(key)
	if env == nil {
		return defaultValue
	}

	v, err := strconv.ParseFloat(*env, 64)
	if err != nil {
		e.l.Printf("[INFO] %s is set but is not a float64, using default value\n", key)
		return defaultValue
	}

	return v
}

func (e *Parser) Bool(key string, defaultValue bool) bool {
	env := e.loadEnv(key)
	if env == nil {
		return defaultValue
	}

	v, err := strconv.ParseBool(*env)
	if err != nil {
		e.l.Printf("[INFO] %s is set but is not a boolean, using default value\n", key)
		return defaultValue
	}

	return v
}

func (e *Parser) loadEnv(key string) *string {
	env, ok := os.LookupEnv(key)
	if !ok {
		e.l.Printf("[INFO] %s not set, using default value\n", key)
		return nil
	}

	if len(env) == 0 {
		e.l.Printf("[INFO] %s is set but empty, using default value\n", key)
		return nil
	}

	return &env
}
