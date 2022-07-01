package util

import "sync"

func WaitGroup(fns ...func()) {
	if len(fns) == 0 {
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(fns))
	for _, fn := range fns {
		go func(fn func()) {
			defer wg.Done()
			fn()
		}(fn)
	}
	wg.Wait()
}
