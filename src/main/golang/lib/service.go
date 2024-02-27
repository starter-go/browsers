package lib

import (
	"fmt"

	"github.com/starter-go/afs"
	"github.com/starter-go/base/util"
	"github.com/starter-go/browsers"
	"github.com/starter-go/vlog"
)

// BrowserServiceImpl ...
type BrowserServiceImpl struct {

	//starter:component
	_as func(browsers.Service) //starter:as("#")

	Regs []browsers.Registry //starter:inject(".")

	cache []*browsers.Registration
}

func (inst *BrowserServiceImpl) _impl() browsers.Service {
	return inst
}

func (inst *BrowserServiceImpl) loadBrowsers() []*browsers.Registration {
	src := inst.Regs
	dst := make([]*browsers.Registration, 0)
	for _, r1 := range src {
		r2 := r1.Registration()
		if r2.Enabled && r2.Browser != nil {
			dst = append(dst, r2)
		}
	}
	s := &util.Sorter{
		OnLen:  func() int { return len(dst) },
		OnLess: func(i1, i2 int) bool { return dst[i1].Priority > dst[i2].Priority },
		OnSwap: func(i1, i2 int) { dst[i1], dst[i2] = dst[i2], dst[i1] },
	}
	s.Sort()
	return dst
}

func (inst *BrowserServiceImpl) getBrowsers() []*browsers.Registration {
	all := inst.cache
	if all == nil {
		all = inst.loadBrowsers()
		inst.cache = all
	}
	return all
}

// Run ...
func (inst *BrowserServiceImpl) Run(intent *browsers.Intent) error {
	t, err := inst.Open(intent)
	if err != nil {
		return err
	}
	return t.Run()
}

// Start ...
func (inst *BrowserServiceImpl) Start(intent *browsers.Intent) (browsers.Task, error) {
	t, err := inst.Open(intent)
	if err != nil {
		return nil, err
	}
	err = t.Start()
	if err != nil {
		return nil, err
	}
	return t, nil
}

// Open ...
func (inst *BrowserServiceImpl) Open(intent *browsers.Intent) (browsers.Task, error) {
	all := inst.getBrowsers()
	for _, item := range all {
		b := item.Browser
		if b.Accept(intent) {
			return b.Open(intent)
		}
	}
	return nil, fmt.Errorf("no browser can handle the intent")
}

func locateExecutable(exe afs.Path, browser string) afs.Path {

	if exe != nil {
		if exe.IsFile() {
			return exe
		}
	}

	// load from default config file
	locator := new(defaultBrowserLocator)
	path, err := locator.locate(browser)
	if err != nil {
		vlog.Warn(err.Error())
		return exe
	}

	return path
}
