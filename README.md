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

    name := p.String("NAME", "Csothen")

    age := p.Int("AGE", 22)
    height := p.Int32("HEIGHT", 171)
    port := p.Int64("PORT", 8080)

    pi := p.Float32("PI", 3.14)
    basePrice := p.Float64("BASE_PRICE", 100.00)

    show := p.Bool("SHOW_PRICES", false)
}
```
