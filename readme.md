# friendly-public-transport-format

This is a implementation
of [fptf](https://github.com/public-transport/friendly-public-transport-format/blob/1.2.1/spec/readme.md)
in golang.

Its just a collection of types, really. But since fptf requires some optional

## Usage

Import the stuff:

```go
package main

import (
	"encoding/json"
	"github.com/Vector-Hector/friendly-public-transport-format"
)

func main() {
	// ...find some data called dat

	var journey fptf.Journey
	err := json.Unmarshal(dat, &journey)
	if err != nil {
		panic(err)
	}
}
```

All the types specified in the [specs](https://github.com/public-transport/friendly-public-transport-format/blob/1.2.1/spec/readme.md)
can be accessed through this package. \
Just capitalize the first letter and write fptf. in front of it ^^