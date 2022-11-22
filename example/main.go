package main

import (
	"fmt"
	"github.com/golang-infrastructure/go-singleflight"
	"math/rand"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	g := singleflight.Group[int]{}

	for i := 0; i < 10; i++ {
		id := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			v, err, shared := g.Do("test", func() (value int, err error) {
				fmt.Println(id, "开始执行了...")
				time.Sleep(time.Second * 3)
				return rand.Intn(100), nil
			})
			fmt.Println(id, v, err, shared)
		}()
	}
	wg.Wait()
	fmt.Println("All done")

}
