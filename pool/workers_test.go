package pool

import (
	"fmt"
	"sync"
	"testing"
)

func TestWorkPool_Start(t *testing.T) {
	wg := sync.WaitGroup{}
	wp := NewPool(2, 50).Start()
	lenth := 100
	wg.Add(lenth)
	for i := 0; i < lenth; i++ {
		wp.PushTaskFunc(func(args ...interface{}) {
			defer wg.Done()
			fmt.Print(args[0].(int), " ")

		}, i)
	}
	wg.Wait()
}
