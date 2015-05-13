# bspc-go

[bspwm](http://github.com/baskerville/bspwm) is a tiling window manager, where
monitor/desktop/window manipulation occurs using the command-line tool `bspc`. bspc-go is a Golang wrapper for this interface, allowing the user to interact with the window manager using Monitor, Desktop, and Window structs

## TODO

* Add usage instructions

## Example

Here is an example of i3-like behavior:

```Go
package main

import (
	"bspc"
	"os"
)

func main() {
	conn, err := bspc.NewController()

	if err != nil {
		return
	}

	desktopName := os.Args[1]

	monitor, _ := conn.FocusedMonitor()

	desktop, _ := monitor.DesktopByName(desktopName)

	if desktop == nil {
		newDesktops, _ := monitor.AddDesktops(desktopName)
		newDesktops[0].Focus()

	} else {
		// Otherwise, we should go to it
		desktop.Focus()
	}

	for _, monitor := range conn.Monitors {
		for _, desktop := range monitor.Desktops {
			if len(desktop.Windows) == 0 {
				desktop.Remove()
			}
		}
	}
}
```
