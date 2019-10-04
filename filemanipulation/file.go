package main

import (
	"io/ioutil"
)

func main() {
	err := ioutil.WriteFile("test.txt", []byte("Hi\n"), 0666)
	if err != nil {
		panic(err)
	}
}
