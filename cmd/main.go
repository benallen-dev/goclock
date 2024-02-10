package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/benallen-dev/goclock/pkg/pretty"
	"github.com/benallen-dev/goclock/pkg/terminal"
)

func printTime (time time.Time) {
	terminal.ResetCursor()

	formatted := pretty.FormatTime(time.Hour(), time.Minute(), -1)

	// Add a little space around the time
	fmt.Println("")

	for _, line := range formatted {
		fmt.Println(" ", line)
	}
}

func main() {
	
	// Clear the console and hide the cursor
	terminal.ClearConsole()
	terminal.HideCursor()

	// Print the time once before the ticker starts
	printTime(time.Now())
	
	// Start the ticker
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Received done signal. Exiting the goroutine.")
				return
			case t := <-ticker.C:
				printTime(t.Local())
			}

		}
	}()

	// Wait for kill signal
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, os.Interrupt, syscall.SIGTERM)
	
	<-osSignal

	// Show the cursor and clear the console
	terminal.ShowCursor()
	terminal.ClearConsole()
	
	fmt.Println("\nReceived kill signal. Exiting the program.")

	ticker.Stop()
	done <- true

	// There's a better way but I'm lazy so we're just going to sleep for a bit
	time.Sleep(5 * time.Millisecond)

}
