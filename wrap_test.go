package errorpanic

import (
	"fmt"
	"testing"
)

func TestWrapNonPanic(t *testing.T) {
	var err = fmt.Errorf("error")
	f := func() error { return err }
	if Wrap(f) != err {
		t.Fatal("wrapper should return regular error of function")
	}
}

func TestWrapPanicString(t *testing.T) {
	var err = fmt.Errorf("error")
	f := func() error {
		panic("doomed")
		return err
	}
	result := Wrap(f)
	if result.Error() != "doomed" {
		t.Fatal("wrapper should return panic string error")
	}
}

func TestWrapPanicError(t *testing.T) {
	var err = fmt.Errorf("error")
	f := func() error {
		panic(err)
		return err
	}
	result := Wrap(f)
	if result != err {
		t.Fatal("wrapper should return panic error")
	}
}
