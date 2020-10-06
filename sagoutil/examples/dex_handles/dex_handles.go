package main

import (
	"fmt"

	"github.com/Meshbits/shurli-server/sagoutil"
)

func main() {

	var handles []sagoutil.DEXHandle
	handles = sagoutil.DEXHandles()

	fmt.Println(handles)

}
