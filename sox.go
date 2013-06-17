package sox

//#include <sox.h>
//#cgo LDFLAGS: -lsox
import "C"

//import "unsafe"

import "fmt"

/*
Find and load format handler plugins.
*/
func FormatInit() int {
	return int(C.sox_format_init())
}

/*
Unload format handler plugins.
*/
func FormatQuit() {
	C.sox_format_quit()
}

/*
Initialize effects library.
*/
func Init() int {
	return int(C.sox_init())
}

/*
Close effects library and unload format handler plugins.
*/
func Quit() int {
	return int(C.sox_quit())
}

/*
Returns the table of format handler names and functions.
*/
func FormatFns() (hurr []string) {
	var format_tabs *C.sox_format_tab_t

	format_tabs = C.sox_get_format_fns()
  fmt.Printf("%#v\n", format_tabs)
  fmt.Printf("%#v\n", format_tabs.name)

	return
}

type FormatTab struct {
  Name string
  fn sox_format_fn_t 
