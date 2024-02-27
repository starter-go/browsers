package lib

import (
	"github.com/starter-go/afs"
	"github.com/starter-go/browsers"
)

// Firefox ...
type Firefox struct {

	//starter:component
	_as func(browsers.Registry) //starter:as(".")

	FS afs.FS //starter:inject("#")

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

func (inst *Firefox) getPath() afs.Path {
	exe := inst.FS.NewPath(inst.ExePath)
	return locateExecutable(exe, "firefox")
}

// Accept ...
func (inst *Firefox) Accept(i *browsers.Intent) bool {
	exe := inst.getPath()
	return exe.IsFile()
}

// Open  ...
func (inst *Firefox) Open(i *browsers.Intent) (browsers.Task, error) {
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
