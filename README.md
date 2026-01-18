# aliasgen

Small zero dependency Go library for generating **cryptographically secure** random aliases/strings.

It uses `crypto/rand` (CSPRNG) and an URL-friendly alphabet:

```
abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_.~
```

## Install

```bash
go get -u github.com/withoutasecondthought/aliasgen
```

## Usage

### Generate a random string (recommended)

```go
package main

import (
	"fmt"

	"aliasgen"
)

func main() {
	// Default length is 5
	s, err := aliasgen.GenerateRandomString()
	if err != nil {
		// Very rare: OS CSPRNG read failed
		panic(err)
	}
	fmt.Println(s)

	// Custom length
	long, err := aliasgen.GenerateRandomString(10)
	if err != nil {
		panic(err)
	}
	fmt.Println(long)
}
```

### Panic on failure

If you prefer a convenience helper that panics on error:

```go
s := aliasgen.MustGenerateRandomString()
long := aliasgen.MustGenerateRandomString(32)
```

## API

- `GenerateRandomString(l ...int) (string, error)`
  - If no length is provided, defaults to `5`.
  - If a non-positive length is provided (`0` or negative), it falls back to the default.
  - Returns a wrapped error with `ErrGenerateRandomString` when random generation fails.

- `MustGenerateRandomString(l ...int) string`
  - Calls `GenerateRandomString` and panics if it returns an error.

## Error handling

When generation fails, `GenerateRandomString` returns an error that wraps:

- `aliasgen.ErrGenerateRandomString` (package sentinel error)
- the underlying OS/reader error from `crypto/rand`

Example:

```go
s, err := aliasgen.GenerateRandomString()
if err != nil {
	if errors.Is(err, aliasgen.ErrGenerateRandomString) {
		// handle generation failure
	}
}
```

## License MIT
