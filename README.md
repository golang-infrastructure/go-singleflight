# go singleflight 泛型版

# 一、这是啥

Go官方扩展库singleflight的泛型实现：

```text
https://cs.opensource.google/go/x/sync
```



# 二、安装

```go
go get -u github.com/golang-infrastructure/go-singleflight
```



# 三、Example

```go
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

	begin := time.Now()
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
	fmt.Println("cost: ", time.Now().Sub(begin).String())
	fmt.Println("All done")
}
```



# 四、TODO

- 写篇文章讨论一下singleflight，分析一下源码
- 跑下benchmark，看下泛型版对性能到底有多大的影响 



