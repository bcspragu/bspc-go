package bspc

import (
	"bytes"
	"errors"
	"os/exec"
)

type Controller struct {
	Monitors []*Monitor
}

func NewController() (*Controller, error) {
	c := new(Controller)

	monitors, err := Monitors()
	if err != nil {
		return nil, err
	}
	c.Monitors = monitors
	return c, nil
}

func (c *Controller) MonitorByIndex(index int) (*Monitor, error) {
	if index < 0 {
		return nil, errors.New("Need to enter a non-negative index")
	}

	if index >= len(c.Monitors) {
		return nil, errors.New("Invalid index")
	}

	return c.Monitors[index], nil
}

func (c *Controller) MonitorByName(name string) (*Monitor, error) {
	for _, monitor := range c.Monitors {
		if monitor.Name == name {
			return monitor, nil
		}
	}
	return nil, errors.New("Monitor not found")
}

func (c *Controller) DesktopByName(name string) (*Desktop, error) {
	for _, monitor := range c.Monitors {
		for _, desktop := range monitor.Desktops {
			if desktop.Name == name {
				return desktop, nil
			}
		}
	}
	return nil, errors.New("Desktop not found")
}

func (c *Controller) WindowByID(id string) (*Window, error) {
	for _, monitor := range c.Monitors {
		for _, desktop := range monitor.Desktops {
			for _, window := range desktop.Windows {
				if window.ID == id {
					return window, nil
				}
			}
		}
	}
	return nil, errors.New("Window not found")

}

func (c *Controller) FocusedMonitor() (*Monitor, error) {
	out, err := exec.Command(CommandName, "query", "--monitor", "focused", "--monitors").Output()
	out = bytes.TrimSpace(out)
	if err != nil {
		return nil, err
	}

	return c.MonitorByName(string(out))
}

func (c *Controller) RemoveEmptyDesktops() error {
	for _, monitor := range c.Monitors {
		err := monitor.RemoveEmptyDesktops()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Controller) Defrag() error {
	for _, monitor := range c.Monitors {
		err := monitor.DefragDesktops()
		if err != nil {
			return err
		}
	}
	return nil
}
