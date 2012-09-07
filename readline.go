package readline

// #cgo LDFLAGS: -lreadline
// #include <stdio.h>
// #include <stdlib.h>
// #include <readline/readline.h>
import "C"
import (
	"unsafe"
)

func init() {
	C.rl_initialize()
}

func Readline(prompt string) string {
	ptr := C.CString(prompt)
	defer free(ptr)
	return C.GoString(C.readline(C.CString(prompt)))
}

func free(v *C.char) {
	C.free(unsafe.Pointer(v))
}
