# go singleflight 泛型版

# 一、这是啥

Go官方扩展库singleflight的泛型实现：

```go
https: //pkg.go.dev/golang.org/x/sync/singleflight
```

# 二、Example

```go
package main

import (
	"fmt"
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

```



