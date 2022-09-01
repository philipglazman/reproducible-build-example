package main

// #include <stdio.h>
// #include <stdlib.h>
//
// static void print(char* s) {
//   printf("%s\n", s);
// }
import "C"

import (
	"unsafe"
)

func main() {
	v := C.CString("Hello from the C Heap!")
	C.print(v)
	C.free(unsafe.Pointer(v))
}
