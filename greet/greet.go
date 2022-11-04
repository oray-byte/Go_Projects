package main

import (
	"fmt"

	"github.com/oray-byte/greeting/greeting"
)

func main() {
	message := greeting.Hello("Owen")

	fmt.Println(message)
}
