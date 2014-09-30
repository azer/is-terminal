package isterminal

import (
	"fmt"
	"syscall"
)

func Example_Test () {
	fmt.Println(IsTerminal(syscall.Stdout))
	// Output: true
}
