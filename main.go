package main

import (
	"os"

	"github.com/borsTiHD/go-toastify/cli"
	"github.com/borsTiHD/go-toastify/gui"
)

func main() {
	// If no arguments provided (or only the executable name), launch GUI
	if len(os.Args) < 2 {
		gui.Run()
		return
	}

	// Otherwise, handle CLI commands
	if cli.Run(os.Args) {
		return
	}

	// Fallback to GUI if no valid CLI command was processed
	gui.Run()
}
