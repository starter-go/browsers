package lib

import (
	"fmt"
	"os"
	"os/user"

	"github.com/starter-go/afs"
	"github.com/starter-go/afs/files"
	"github.com/starter-go/application/properties"
)

type defaultBrowserLocator struct {
}

func (inst *defaultBrowserLocator) locate(name string) (afs.Path, error) {

	cfg, err := inst.getConfigFile()
	if err != nil {
		return nil, err
	}

	if !cfg.Exists() {
		inst.createConfigFile(cfg)
	}

	props, err := inst.loadConfigProps(cfg)
	if err != nil {
		return nil, err
	}

	key := "browser." + name + ".executable"
	value, err := props.GetPropertyRequired(key)
	if err != nil {
		return nil, fmt.Errorf("no property [%s] in file [%s], %s", key, cfg.GetPath(), err)
	}

	exe := cfg.GetFS().NewPath(value)
	if !exe.IsFile() {
		return nil, fmt.Errorf("no executable file at path [%s]", exe.GetPath())
	}

	return exe, nil
}

func (inst *defaultBrowserLocator) createConfigFile(cfg afs.Path) error {
	cfg.MakeParents(&afs.Options{
		Permission: 0644,
	})
	opt := &afs.Options{
		Flag:       os.O_CREATE | os.O_WRONLY,
		Permission: 0644,
	}
	return cfg.GetIO().WriteText("", opt)
}

func (inst *defaultBrowserLocator) loadConfigProps(file afs.Path) (properties.Table, error) {
	txt, err := file.GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}
	return properties.Parse(txt, nil)
}

func (inst *defaultBrowserLocator) getConfigFile() (afs.Path, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	home := files.FS().NewPath(u.HomeDir)
	cfg := home.GetChild(".browsers/browsers.config")
	return cfg, nil
}
