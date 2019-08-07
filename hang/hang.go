package hang

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func HangOn(clear func()) {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	sig := <-sc
	switch sig {
	case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
		fmt.Println("wait to clear work environment!")
		clear()
		fmt.Println("service shutdown now!")
		os.Exit(0)
	}
}
