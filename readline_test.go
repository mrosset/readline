package readline

import (
	"fmt"
	"testing"
)

func TestReadlin(t *testing.T) {
	f := Readline(">")
	fmt.Println(f)
}
