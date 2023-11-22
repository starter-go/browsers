package unit

import (
	"github.com/starter-go/application"
	"github.com/starter-go/browsers"
)

// Test1 ...
type Test1 struct {
	//starter:component
	//// _as func(application.Lifecycle) //starter:as(".")

	Browsers browsers.Service //starter:inject("#")
}

func (inst *Test1) _impl() application.Lifecycle {
	return inst
}

// Life ...
func (inst *Test1) Life() *application.Life {
	return &application.Life{
		OnStart: inst.start,
	}
}

func (inst *Test1) start() error {
	i := &browsers.Intent{
		URL: "https://www.bitwormhole.com/demo/foo/bar",
	}
	return inst.Browsers.Run(i)
}
