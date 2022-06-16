# Go library for weatherapi.com

### Easy-to-use api weather.com

## Installing

` go get github.com/Bogdan-Torkhov/weatherapi-go `

## Using

```go
package main
import (
	"fmt"
	weather "github.com/Bogdan-Torkhov/weatherapi-go"
)
func main() {
	const yourAPIkey = "apiKeyHere"
	w, err := weather.GetWeather(yourAPIkey, "London") // Api key, location
	if err != nil {
		panic(err)
	}
	fmt.Println(w) // All weather
	fmt.Println(w.GetMain()) // Only main weather
	fmt.Println(w.Current.Condition.Text) // Output: sunny
}
```
