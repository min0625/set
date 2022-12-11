# Go Set
[![Go Reference](https://pkg.go.dev/badge/github.com/min0625/set.svg)](https://pkg.go.dev/github.com/min0625/set)

This is a `Set` based on `Generic` and `Map` implementations.

## Installation
```sh
go get github.com/min0625/set
```

## Quick start
```go
package main

import (
	"fmt"

	"github.com/min0625/set"
)

func main() {
	s := make(set.Set[int])

	s.Store(1, 2, 3)
	fmt.Println(s.Has(1))

	s.Delete(1)
	fmt.Println(s.Has(1))

	// Output:
	// true
	// false
}
```

## Example
See: [./set_example_test.go](./set_example_test.go)
