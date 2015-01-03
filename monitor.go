package bspc

import (
	"bytes"
	"os/exec"
)

type Monitor struct {
	Name     string
	Desktops []*Desktop
}

func Monitors() ([]*Monitor, error) {
	out, err := exec.Command(CommandName, "query", "--monitors").Output()
	if err != nil {
		return nil, err
	}
	names := bytes.Split(out, []byte("\n"))
	// Last entry is blank
	monitors := make([]*Monitor, len(names)-1)

	for i, name := range names[:len(names)-1] {
		monitors[i] = &Monitor{Name: string(name)}
		err = monitors[i].loadDesktops()
		if err != nil {
			return nil, err
		}
	}
	return monitors, nil
}

func (m *Monitor) loadDesktops() error {
	out, err := exec.Command(CommandName, "query", "--monitor", m.Name, "--desktops").Output()

	if err != nil {
		return err
	}

	names := bytes.Split(out, []byte("\n"))
	// Last entry is blank
	desktops := make([]*Desktop, len(names)-1)

	for i, name := range names[:len(names)-1] {
		desktops[i] = &Desktop{Name: string(name)}
		err = desktops[i].loadWindows()
		if err != nil {
			return err
		}
	}
	m.Desktops = desktops
	return nil
}

func (m *Monitor) Rename(newName string) error {
	err := exec.Command(CommandName, "monitor", m.Name, "--rename", newName).Run()
	if err != nil {
		return err
	}
	m.Name = newName
	return nil
}

func (m *Monitor) AddDesktops(desktops ...string) error {
	for _, desktop := range desktops {
		err := exec.Command(CommandName, "monitor", m.Name, "--add-desktops", desktop).Run()
		if err != nil {
			return err
		}
		m.Desktops = append(m.Desktops, &Desktop{Name: desktop})
	}
	return nil
}
