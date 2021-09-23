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

type EnvParser interface {
	Load(file string) error
	String(key, defaultValue string) string
	Int(key string, defaultValue int) int
	Int32(key string, defaultValue int32) int32
	Int64(key string, defaultValue int64) int64
	Float32(key string, defaultValue float32) float32
	Float64(key string, defaultValue float64) float64
	Bool(key string, defaultValue bool) bool
}

type Parser struct {
	cache map[string]string
}

// NewParser creates a new Parser instance
func NewParser() *Parser {
	return &Parser{
		cache: make(map[string]string),
	}
}

// Load loads all the environment variables
// present in the file passed
func (p *Parser) Load(file string) error {
	f, err := os.Open(file)

	if err != nil {
		return fmt.Errorf("could not open the file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "=")
		if len(line) != 2 {
			return ErrWrongFormat
		}

		key, value := strings.TrimSpace(line[0]), strings.TrimSpace(line[1])

		p.cache[key] = value

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

func (p *Parser) String(key, defaultValue string) string {
	v, isCached := p.cache[key]
	if !isCached {
		env := loadEnv(key)
		if env == nil {
			return defaultValue
		}
		return *env
	}
	return v
}

func (p *Parser) Int(key string, defaultValue int) int {
	v, isCached := p.cache[key]
	if !isCached {
		env := loadEnv(key)
		if env == nil {
			return defaultValue
		}

		return toInt(*env, defaultValue)
	}

	return toInt(v, defaultValue)
}

func (p *Parser) Int32(key string, defaultValue int32) int32 {
	v, isCached := p.cache[key]
	if !isCached {
		env := loadEnv(key)
		if env == nil {
			return defaultValue
		}

		return toInt32(*env, defaultValue)
	}

	return toInt32(v, defaultValue)
}

func (p *Parser) Int64(key string, defaultValue int64) int64 {
	v, isCached := p.cache[key]
	if !isCached {
		env := loadEnv(key)
		if env == nil {
			return defaultValue
		}

		return toInt64(*env, defaultValue)
	}

	return toInt64(v, defaultValue)
}

func (p *Parser) Float32(key string, defaultValue float32) float32 {
	v, isCached := p.cache[key]
	if !isCached {
		env := loadEnv(key)
		if env == nil {
			return defaultValue
		}

		return toFloat32(*env, defaultValue)
	}

	return toFloat32(v, defaultValue)
}

func (p *Parser) Float64(key string, defaultValue float64) float64 {
	v, isCached := p.cache[key]
	if !isCached {
		env := loadEnv(key)
		if env == nil {
			return defaultValue
		}

		return toFloat64(*env, defaultValue)
	}

	return toFloat64(v, defaultValue)
}

func (p *Parser) Bool(key string, defaultValue bool) bool {
	v, isCached := p.cache[key]
	if !isCached {
		env := loadEnv(key)
		if env == nil {
			return defaultValue
		}

		return toBool(*env, defaultValue)
	}

	return toBool(v, defaultValue)
}
