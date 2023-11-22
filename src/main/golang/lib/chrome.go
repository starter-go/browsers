package lib

import "github.com/starter-go/browsers"

// Chrome ...
type Chrome struct {

	//starter:component
	_as func(browsers.Registry) //starter:as(".")

	ExePath  string //starter:inject("${browser.chrome.executable}")
	Enabled  bool   //starter:inject("${browser.chrome.enabled}")
	Priority int    //starter:inject("${browser.chrome.priority}")
}

func (inst *Chrome) _impl() (browsers.Browser, browsers.Registry) {
	return inst, inst
}

// Registration ...
func (inst *Chrome) Registration() *browsers.Registration {
	return &browsers.Registration{
		Enabled:  inst.Enabled,
		Priority: inst.Priority,
		Browser:  inst,
	}
}

// Accept ...
func (inst *Chrome) Accept(i *browsers.Intent) bool {
	exe := getPath(inst.ExePath)
	return exe.IsFile()
}

// Open ...
func (inst *Chrome) Open(i *browsers.Intent) (browsers.Task, error) {
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
