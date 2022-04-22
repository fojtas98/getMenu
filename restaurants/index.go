package restaurants

import (
	"fmt"
	"sync"
)

func Fetch() {
	c := make(chan string, 3)
	wg := sync.WaitGroup{}
	wg.Add(3)

	go uTelleru(c, &wg)
	go letsmeat(c, &wg)
	go vinPivovar(c, &wg)
	wg.Wait()
	close(c)
	for v := range c {
		fmt.Println(v)
	}

}
