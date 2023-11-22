package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

	"github.com/starter-go/afs"
	"github.com/starter-go/browsers"
	"github.com/starter-go/vlog"
)

type task struct {
	Context    context.Context
	Executable afs.Path
	Args       []string

	starting bool
	started  bool
	stopping bool
	stopped  bool
	cmd      *exec.Cmd
}

func (inst *task) handleError(err error) {
	if err == nil {
		return
	}
	vlog.Error(err.Error())
}

func (inst *task) handleErrorX(x any) {

	if x == nil {
		return
	}

	err, ok := x.(error)
	if ok {
		inst.handleError(err)
		return
	}

	str, ok := x.(string)
	if ok {
		inst.handleError(fmt.Errorf(str))
		return
	}

	data, err := json.Marshal(x)
	if err != nil {
		inst.handleError(err)
	} else {
		str = string(data)
		inst.handleError(fmt.Errorf(str))
	}
}

func (inst *task) open() (browsers.Task, error) {
	return inst, nil
}

func (inst *task) Start() error {
	if inst.starting {
		return fmt.Errorf("call start only once")
	}
	inst.starting = true
	go func() {
		err := inst.Run()
		inst.handleError(err)
	}()
	return inst.waitForStarted()
}

func (inst *task) Run() error {

	defer func() {
		inst.stopped = true
		x := recover()
		inst.handleErrorX(x)
	}()

	ctx := inst.Context
	name := inst.Executable.GetPath()
	args := inst.Args

	if ctx == nil {
		ctx = context.Background()
	}

	cmd := exec.CommandContext(ctx, name, args...)
	inst.cmd = cmd
	inst.started = true

	// err := cmd.Run()
	err := cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Wait()
}

func (inst *task) Join() error {
	cmd := inst.cmd
	if cmd == nil {
		return fmt.Errorf("task is not started")
	}
	return cmd.Wait()
}

func (inst *task) waitForStarted() error {
	if !inst.starting {
		return fmt.Errorf("call start first")
	}
	for {
		if inst.stopped {
			return fmt.Errorf("task has stopped")
		}
		if inst.started {
			return nil
		}
		time.Sleep(time.Millisecond * 100)
	}
}
