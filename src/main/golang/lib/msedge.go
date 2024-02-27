package lib

import (
	"github.com/starter-go/afs"
	"github.com/starter-go/browsers"
)

// MSEdge ...
type MSEdge struct {

	//starter:component
	_as func(browsers.Registry) //starter:as(".")

	FS afs.FS //starter:inject("#")

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

func (inst *MSEdge) getPath() afs.Path {
	exe := inst.FS.NewPath(inst.ExePath)
	return locateExecutable(exe, "msedge")
}

// Accept ...
func (inst *MSEdge) Accept(i *browsers.Intent) bool {
	exe := inst.getPath()
	return exe.IsFile()
}

// Open ...
func (inst *MSEdge) Open(i *browsers.Intent) (browsers.Task, error) {
	url := i.URL
	exe := inst.getPath()
	args := []string{url}
	t := &task{
		Context:    i.Context,
		Executable: exe,
		Args:       args,
	}
	return t.open()
}
