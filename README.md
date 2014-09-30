## is-terminal

Returns true or false depending on if the `os.Stdout` is associated with a terminal.

## Install

```bash
$ go get github.com/azer/is-terminal
```

## Usage

```go
package main

import (
  "github.com/azer/is-terminal"
  "fmt"
)

func () {
  fmt.Println(isterminal.IsTerminal())
  // => true or false
}
```
