package lib

import "github.com/starter-go/browsers"

// MSEdge ...
type MSEdge struct {

	//starter:component
	_as func(browsers.Registry) //starter:as(".")

	ExePath  string //starter:inject("${browser.msedge.executable}")
	Enabled  bool   //starter:inject("${browser.msedge.enabled}")
	Priority int    //starter:inject("${browser.msedge.priority}")
}

func (inst *MSEdge) _impl() (browsers.Browser, browsers.Registry) {
	return inst, inst
}

// Registration ...
func (inst *MSEdge) Registration() *browsers.Registration {
	return &browsers.Registration{
		Enabled:  inst.Enabled,
		Priority: inst.Priority,
		Browser:  inst,
	}
}

// Accept ...
func (inst *MSEdge) Accept(i *browsers.Intent) bool {
	exe := getPath(inst.ExePath)
	return exe.IsFile()
}

// Open ...
func (inst *MSEdge) Open(i *browsers.Intent) (browsers.Task, error) {
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
