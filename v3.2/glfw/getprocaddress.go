package glfw

//#include "glfw/include/GLFW/glfw3.h"
//#include <stdlib.h>
import "C"

import (
	"unsafe"
)

// GetProcAddress provides an exported wrapper around
// the C.glfwGetProcAddress function.
func GetProcAddress(procname string) unsafe.Pointer {
	cProcname := C.CString(procname)
	defer C.free(unsafe.Pointer(cProcname))

	return unsafe.Pointer(C.glfwGetProcAddress(cProcname))
}
