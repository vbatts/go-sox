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
Search for effect by name. nil is returned if nothing found.
*/
func FindEffect(name string) *Effect {
	return newEffect(C.sox_find_effect(C.CString(name)))
}

func newEffect(eh *C.sox_effect_handler_t) *Effect {
	if eh == nil {
		return nil
	}
	return &Effect{
		Name:   C.GoString(eh.name),
		Usage:  C.GoString(eh.usage),
		effect: eh,
	}
}

type Effect struct {
	Name   string
	Usage  string
	Flags  uint
	effect *C.sox_effect_handler_t
}

func (e *Effect) cEffect() (eh *C.sox_effect_handler_t) {
	if e.effect != nil {
		eh = e.effect
	}
	eh.name = C.CString(e.Name)
	eh.usage = C.CString(e.Usage)
	eh.flags = C.uint(e.Flags)
	return eh
}

/*
Returns the table of format handler names and functions.
*/
func formatFns() (hurr []string) {
	var format_tabs *C.sox_format_tab_t

	format_tabs = C.sox_get_format_fns()
	fmt.Printf("%#v\n", format_tabs)
	fmt.Printf("%#v\n", format_tabs.name)

	return
}

type Format struct {
}

func (f *Format) AddEffect(e Effect) error {
	return nil
}

func (f *Format) Read(p []byte) (n int, err error) {
	return
}

func (f *Format) Close() (err error) {
	return nil
}

type formatTab struct {
	Name string
	fn   C.sox_format_fn_t
}
