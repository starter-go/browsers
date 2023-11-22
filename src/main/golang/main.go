package main

import (
	"os"

	"github.com/starter-go/browsers/modules/browsers"
	"github.com/starter-go/starter"
)

func main() {
	m := browsers.Module()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
