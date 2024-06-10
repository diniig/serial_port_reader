package main

import (
	"fmt"
	"log"
	"time"

	"go.bug.st/serial"
)

func main() {
	// Retrieve the port list
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	// Print the list of detected ports
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	// Open the first serial port detected at 9600bps N81
	mode := &serial.Mode{
		// BaudRate: 9600,
		BaudRate: 19200, //?
		// BaudRate: 57600,
		// BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(ports[0], mode)
	if err != nil {
		log.Fatal(err)
	}

	defer port.Close()

	buff := make([]byte, 100)
	for {
		start := time.Now()
		// Reads up to 100 bytes
		n, err := port.Read(buff)
		if time.Now().Sub(start) > time.Second {
			fmt.Printf("==========================\n")
		}
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}

		fmt.Printf("%v\n", buff[:n])

		// If we receive a newline stop reading
		// if strings.Contains(string(buff[:n]), "\n") {
		// 	break
		// }
	}
}
