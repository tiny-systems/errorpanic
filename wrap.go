package errorpanic

import (
	"fmt"
	"sync"
)

func Wrap(f func() error) error {
	var rErr, err error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				switch x := r.(type) {
				case error:
					rErr = x
				default:
					rErr = fmt.Errorf("%v", x)
				}
			}
			wg.Done()
		}()
		err = f()
	}()
	wg.Wait()
	if rErr != nil {
		return rErr
	}
	return err
}
