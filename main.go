package main

import (
	"fmt"
	"github.com/jnst/microservice-x/module1/db"
	"github.com/jnst/microservice-x/module3"
)

func main() {
	fmt.Println(db.Now())
	module3.Hello()
}
