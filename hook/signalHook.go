package devtool

import (
	"os"
	"os/signal"
	"syscall"
)

// CleanupHook return channel:
// Notify os.Interrupt | syscall.SIGTERM
func CleanupHook() (c chan os.Signal) {
	c = make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	return
}
