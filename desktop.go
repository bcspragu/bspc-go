package bspc

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Desktop struct {
	Name    string
	Windows []*Window
}

func NewDesktop(name string, monitor *Monitor) *Desktop {
	desktop := &Desktop{Name: name}
	return desktop
}

func (d *Desktop) Remove() {
	exec.Command(d.getSelector(), "--remove").Run()
}

func (d *Desktop) getSelector() string {
	return fmt.Sprintf("%s desktop %s", CommandName, d.Name)
}

func (d *Desktop) loadWindows() error {
	out, err := exec.Command(CommandName, "query", "--desktop", d.Name, "--windows").Output()

	if err != nil {
		return err
	}

	names := bytes.Split(out, []byte("\n"))
	// Last entry is blank
	windows := make([]*Window, len(names)-1)

	for i, name := range names[:len(names)-1] {
		windows[i] = &Window{ID: string(name)}
	}
	d.Windows = windows
	return nil
}
