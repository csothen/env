package env

import (
	"log"
	"os"
)

func test() {
	l := log.New(os.Stdout, "[ logger ]", log.LstdFlags)
	p := NewParser(l)

	p.Int("PORT", 8080)
	p.Int32("PORT", 8080)
	p.Int64("PORT", 8080)

	p.Float32("BASE_PRICE", 10.00)
	p.Float64("BASE_PRICE", 100.00)

	p.Bool("SHOW_PRICES", false)
}
