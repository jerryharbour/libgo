package osutils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func WaitExit(c chan os.Signal, exit func()) {
	for i := range c {
		switch i {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("receive exit signal ", i.String(), ", exit ...")
			exit()
			os.Exit(0)
		}
	}
}

func NewGracefulexitSignal() chan os.Signal {
	c := make(chan os.Signal)

	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	return c
}
