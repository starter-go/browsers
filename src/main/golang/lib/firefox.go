package lib

import (
	"github.com/starter-go/browsers"
)

// Firefox ...
type Firefox struct {

	//starter:component
	_as func(browsers.Registry) //starter:as(".")

	ExePath  string //starter:inject("${browser.firefox.executable}")
	Enabled  bool   //starter:inject("${browser.firefox.enabled}")
	Priority int    //starter:inject("${browser.firefox.priority}")
}

func (inst *Firefox) _impl() (browsers.Browser, browsers.Registry) {
	return inst, inst
}

// Registration ...
func (inst *Firefox) Registration() *browsers.Registration {
	return &browsers.Registration{
		Enabled:  inst.Enabled,
		Priority: inst.Priority,
		Browser:  inst,
	}
}

// Accept ...
func (inst *Firefox) Accept(i *browsers.Intent) bool {
	exe := getPath(inst.ExePath)
	return exe.IsFile()
}

// Open  ...
func (inst *Firefox) Open(i *browsers.Intent) (browsers.Task, error) {
	url := i.URL
	exe := getPath(inst.ExePath)
	args := []string{url}
	t := &task{
		Context:    i.Context,
		Executable: exe,
		Args:       args,
	}
	return t.open()
}
