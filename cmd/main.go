package main

import (
	"fmt"
	"time"
)

func clearConsole() {
	fmt.Print("\033[H\033[2J")
}

func printTime (time time.Time) {
	fmt.Print("\r\033[K") // Clear the line
	fmt.Print("Time: ", time.Local().Format("15 : 04 : 05"))
}

func main() {
	
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	// Clear the console
	clearConsole()

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("\nReceived done signal. Exiting the goroutine.")
				return
			case t := <-ticker.C:
				time := t.Local()
				// Do the sexy printing in the console
				printTime(time)
			}

		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- true
}
