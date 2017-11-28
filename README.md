# vmt

Simple fmt for processing format like ? and $1

## Getting Started

Sometime you need to expand a query string in golang using question mark like this

`SELECT * from user where age=? and gender=?` with param=(25, m)

or postgre query like this

`SELECT * from user where age=$1 and gender=$2` with param=(25, m)

using this library you can just use like this

`fmt.Println(vmt.Sprintq(`SELECT * from user where age=? and gender=?`, 25, "m"))`

or like this

`fmt.Println(vmt.Sprints(`SELECT * from user where age=$1 and gender=$2`, 25, "m"))`


## Complete sample
```go
package main

import (
	"fmt"
	"github.com/arfan/vmt"
)

func main() {
	fmt.Println("hello")

	str:="I $1 $2"

	fmt.Println(vmt.Sprints(str, "love", "you"))

}
```