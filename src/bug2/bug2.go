package bug2

import (
	"fmt"
	"sync"
)

func bug2(n int, foo func()) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			fmt.Println(i)
			foo()
			wg.Done()
		}(i)
	}
	wg.Wait()
}
