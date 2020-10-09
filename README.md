# go-semver

go-semver is simple semver library for Golang.

# Usage

```go
package main

import (
    "fmt"

    "github.com/x-foby/go-semver"
)

func main() {
    v1, _ := semver.Parse("v1.0.0")
    v2, _ := semver.Parse("v1.0.1")

    fmt.Println(v1.Less(v2)) // true
}
```
