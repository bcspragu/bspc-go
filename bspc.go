package bspc

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
