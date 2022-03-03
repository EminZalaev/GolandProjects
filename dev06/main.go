//Реализовать все возможные способы остановки выполнения горутины.
package main

import (
	"sync"
)

// Через канал
// Через контекст

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan int)
	go func() {
		for {
			foo, ok := <-ch
			if !ok {
				println("done")
				wg.Done()
				return
			}
			println(foo)
		}
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	wg.Wait()
}

//  Прокидываем контекст
//func main() {
//	forever := make(chan struct{})
//	ctx, cancel := context.WithCancel(context.Background())
//
//	go func(ctx context.Context) {
//		for {
//			select {
//			case <-ctx.Done():  // if cancel() execute
//				forever <- struct{}{}
//				return
//			default:
//				fmt.Println("for loop")
//			}
//
//			time.Sleep(500 * time.Millisecond)
//		}
//	}(ctx)
//
//	go func() {
//		time.Sleep(3 * time.Second)
//		cancel()
//	}()
//
//	<-forever
//	fmt.Println("finish")
//}
