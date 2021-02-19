# concurrent
 concurrent is a simple and efficient Golang concurrency controller, ``` new ``` is userd to create number of concurrency. ``` add ``` is used to add concurrency to quene. ``` done ``` is used to do a task. ``` wait ``` is used to wait task accomplished.

## Installation
``` 
go get github.com/yunlongliang/concurrent
```
## Import
``` go
import "github.com/yunlongliang/concurrent"
```
## Getting started
``` go
package main

import "github.com/yunlongliang/concurrent"

func main() {
	c := concurrent.New(100)
	for i := 0; i < 1000; i ++ {
		c.Add(1)
		go func(i int) {
			defer c.Done()
			fmt.Println(i)
		}(i)
	}
	c.Wait()
}
```

