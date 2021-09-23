package env

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	l *log.Logger
}

// NewParser creates a new Parser instance
func NewParser(logger *log.Logger) *Parser {
	return &Parser{logger}
}

// Load loads all the environment variables
// present in the file passed
func (p *Parser) Load(file string) error {
	f, err := os.Open(file)

	if err != nil {
		p.l.Printf("[ERROR] Could not open the file %s", file)
		return fmt.Errorf("could not open the file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "=")
		if len(line) != 2 {
			p.l.Printf("[ERROR] The file %s is wrongly formatted", file)
			return ErrWrongFormat
		}
		key, value := line[0], line[1]
		err = os.Setenv(key, value)
		if err != nil {
			p.l.Printf("[ERROR] Could not set the environment variable %s to %s", key, value)
			return fmt.Errorf("could not set environment variable: %v", err)
		}
	}

	if scanner.Err() != nil {
		p.l.Printf("[ERROR] An error occurred while scanning the file %s", file)
		return fmt.Errorf("error while scanning the file: %v", err)
	}

	return nil
}

func (p *Parser) String(key, defaultValue string) string {
	env := p.loadEnv(key)
	if env == nil {
		return defaultValue
	}

	return *env
}

func (p *Parser) Int(key string, defaultValue int) int {
	env := p.loadEnv(key)
	if env == nil {
		return defaultValue
	}

	v, err := strconv.Atoi(*env)
	if err != nil {
		p.l.Printf("[INFO] %s is set but is not an int, using default value\n", key)
		return defaultValue
	}

	return v
}

func (p *Parser) Int32(key string, defaultValue int32) int32 {
	env := p.loadEnv(key)
	if env == nil {
		return defaultValue
	}

	v, err := strconv.ParseInt(*env, 10, 32)
	if err != nil {
		p.l.Printf("[INFO] %s is set but is not an int32, using default value\n", key)
		return defaultValue
	}

	return int32(v)
}

func (p *Parser) Int64(key string, defaultValue int64) int64 {
	env := p.loadEnv(key)
	if env == nil {
		return defaultValue
	}

	v, err := strconv.ParseInt(*env, 10, 64)
	if err != nil {
		p.l.Printf("[INFO] %s is set but is not an int64, using default value\n", key)
		return defaultValue
	}

	return v
}

func (p *Parser) Float32(key string, defaultValue float32) float32 {
	env := p.loadEnv(key)
	if env == nil {
		return defaultValue
	}

	v, err := strconv.ParseFloat(*env, 32)
	if err != nil {
		p.l.Printf("[INFO] %s is set but is not a float32, using default value\n", key)
		return defaultValue
	}

	return float32(v)
}

func (p *Parser) Float64(key string, defaultValue float64) float64 {
	env := p.loadEnv(key)
	if env == nil {
		return defaultValue
	}

	v, err := strconv.ParseFloat(*env, 64)
	if err != nil {
		p.l.Printf("[INFO] %s is set but is not a float64, using default value\n", key)
		return defaultValue
	}

	return v
}

func (p *Parser) Bool(key string, defaultValue bool) bool {
	env := p.loadEnv(key)
	if env == nil {
		return defaultValue
	}

	v, err := strconv.ParseBool(*env)
	if err != nil {
		p.l.Printf("[INFO] %s is set but is not a boolean, using default value\n", key)
		return defaultValue
	}

	return v
}

func (p *Parser) loadEnv(key string) *string {
	env, ok := os.LookupEnv(key)
	if !ok {
		p.l.Printf("[INFO] %s not set, using default value\n", key)
		return nil
	}

	if len(env) == 0 {
		p.l.Printf("[INFO] %s is set but empty, using default value\n", key)
		return nil
	}

	return &env
}
