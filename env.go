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
		if len(line) < 2 {
			return ErrWrongFormat
		}

		key := strings.TrimSpace(line[0])
		value := strings.TrimSpace(strings.Join(line[1:], ""))

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
	v, useDefault := p.envOrDefault(key)
	if useDefault {
		return defaultValue
	}
	return v
}

func (p *Parser) Int(key string, defaultValue int) int {
	v, useDefault := p.envOrDefault(key)
	return toInt(v, defaultValue, useDefault)
}

func (p *Parser) Int32(key string, defaultValue int32) int32 {
	v, useDefault := p.envOrDefault(key)
	return toInt32(v, defaultValue, useDefault)
}

func (p *Parser) Int64(key string, defaultValue int64) int64 {
	v, useDefault := p.envOrDefault(key)
	return toInt64(v, defaultValue, useDefault)
}

func (p *Parser) Float32(key string, defaultValue float32) float32 {
	v, useDefault := p.envOrDefault(key)
	return toFloat32(v, defaultValue, useDefault)
}

func (p *Parser) Float64(key string, defaultValue float64) float64 {
	v, useDefault := p.envOrDefault(key)
	return toFloat64(v, defaultValue, useDefault)
}

func (p *Parser) Bool(key string, defaultValue bool) bool {
	v, useDefault := p.envOrDefault(key)
	return toBool(v, defaultValue, useDefault)
}

func (p *Parser) envOrDefault(key string) (string, bool) {
	v, isCached := p.cache[key]
	if !isCached {
		env := loadEnv(key)
		if env == nil {
			return "", true
		}
		return *env, false
	}
	return v, false
}
