# env

Small library to parse environment variables

## Installation

`go get -u github.com/csothen/env`

## Usage

``` Go
import "github.com/csothen/env"

func main() {
    err := env.Load(".env")
    if err != nil {
        panic(err)
    }

    name := env.String("NAME", "Csothen")

    age := env.Int("AGE", 22)
    height := env.Int32("HEIGHT", 171)
    port := env.Int64("PORT", 8080)

    pi := env.Float32("PI", 3.14)
    basePrice := env.Float64("BASE_PRICE", 100.00)

    show := env.Bool("SHOW_PRICES", false)
}
```
