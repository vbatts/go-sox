package sox

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	i := FormatInit()
	if i != 0 {
		t.Errorf("FormatInit() returned: %d", i)
	}
	defer FormatQuit()
}

func TestFormats(t *testing.T) {
	Init()
	formatFns()
	defer Quit()

}
func TestEffects(t *testing.T) {
	e := FindEffect("vol")
	if e == nil {
		t.Errorf("should not be nil")
	}
	fmt.Printf("%#v\n", e)
}
