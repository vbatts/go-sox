package sox

import (
  //"fmt"
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
  FormatFns()
  defer Quit()

}
