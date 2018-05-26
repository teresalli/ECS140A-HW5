package bug3

import "sync"

// Merge sends the output from channels ch1 and ch2 to out.
func merge(out chan uint, cs ...<-chan uint){
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan uint) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
}


func bug3(producer1 func(chan uint), producer2 func(chan uint)) chan uint {
	// Create channels for producers
	ch1, ch2 := make(chan uint), make(chan uint)
	// Create channel for consumer
	out := make(chan uint)

	// Spawn each goroutine
	go producer1(ch1)
	go producer2(ch2)
	go merge(out, ch1, ch2)

	// Return the output channel
	return out
}
