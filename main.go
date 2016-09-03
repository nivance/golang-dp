package main

import (
	"fmt"

	"github.com/nivance/golang-dp/singleton"
)

func main() {
	s1 := singleton.New()
	fmt.Println("singleton1 create at", s1.CreateTime.String())
	s2 := singleton.New()
	fmt.Println("singleton2 create at", s2.CreateTime.String())
}
