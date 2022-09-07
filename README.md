# tjson
a json  library from golang encoding/json support  custom tag

```go
// 使用tsjon tag，与encoding/json区分
package main

import (
	"fmt"

	tjson "github.com/averyboy/tjson"
)

type A struct {
	A  string `tjson:"名称"`
	ID int64  `tjson:"序号"`
}

func main() {
	a := A{A: "a", ID: 1}
	aa, err := tjson.Marshal(a)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(aa))
}

```