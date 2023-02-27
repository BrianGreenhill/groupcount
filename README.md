# Group Count

Group Count is a waitgroup that exposes its atomic counter.

## Usage

```golang
package main

import (
    "fmt"
    "time"

    "github.com/briangreenhill/groupcount"
)

func main() {
    var wg groupcount.WaitGroup

    fmt.Println("init:", wg.Count())
    wg.Add(1)
    fmt.Println("before:", wg.Count())
    go func() {
        defer wg.Done()
        // do stuff
        time.Sleep(time.Second * 5)
    }()
    wg.Wait()
    fmt.Println("after:", wg.Count())
}
```
