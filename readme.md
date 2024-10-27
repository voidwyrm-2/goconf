# Goconf
Goconf is a library to load and save a simple configuration format I created in 3 minutes

**Example**

```
. config.txt

name:string:Jacob Thaumiel;
age:int: 21;

. since semicolons are used to end fields, use `\;` to escape them
. and if needed, use `\:` to escape colons
```

```go
// main.go
package main

import (
    "fmt"
    "github.com/voidwyrm-2/goconf"
)

func main() {
    config, err := goconf.Load("config.txt")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    // config is a string -> any map

    fmt.Println(config["name"]) // "Jacob Thaumiel"
    fmt.Println(config["age"]) // 21
}
```

All supported Go types can be seen in [goconf.go](./goconf.go)
