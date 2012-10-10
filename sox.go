package sox

//#include <sox.h>
//#cgo LDFLAGS: -lsox
import "C"
//import "unsafe"

func FormatInit() int {
	return int(C.sox_format_init())
}

func FormatQuit() {
	C.sox_format_quit()
}

