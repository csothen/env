# env

Small library to parse environment variables

## Installation

`go get -u github.com/csothen/env`

## Usage

``` Go
import "github.com/csothen/env"

func main() {
    l := log.New(os.Stdout, "[ logger ]", log.LstdFlags)
    p := env.NewParser()

    p.String("DOMAIN", "www.    csothen.com")

    p.Int("PORT", 8080)
    p.Int32("PORT", 8080)
    p.Int64("PORT", 8080)

    p.Float32("BASE_PRICE", 10. 00)
    p.Float64("BASE_PRICE", 100.    00)

    p.Bool("SHOW_PRICES", false)
}
```
