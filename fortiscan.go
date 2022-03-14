package main

import (
	"github.com/lateralusd/fortiscan/scan"
	"io/ioutil"
	"log"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

func main() {
	scan.FromStdin()
}
