package main

import (
	"fmt"
	"time"

	"github.com/rfyiamcool/cronlib"
)

func main() {
	t, err := cronlib.Parse("0 0 0 */1 * *")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(" now: ", time.Now())
	next := t.Next(time.Now())
	fmt.Println("next: ", next)
}

