package main

import (
	"fmt"
	"math"
	"net"
	"os"
	"time"
)

const (
	Red           = "\x1b[31m"
	Green         = "\x1b[32m"
	Yellow        = "\x1b[33m"
	Blue          = "\x1b[34m"
	Magenta       = "\x1b[35m"
	Cyan          = "\x1b[36m"
	White         = "\x1b[37m"
	Reset         = "\x1b[0m"
	Black         = "\x1b[30m"
	BrightRed     = "\x1b[91m"
	BrightGreen   = "\x1b[92m"
	BrightYellow  = "\x1b[93m"
	BrightBlue    = "\x1b[94m"
	BrightMagenta = "\x1b[95m"
	BrightCyan    = "\x1b[96m"
	BrightWhite   = "\x1b[97m"
)

func papingTCP(address string, port string, timeout time.Duration) {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", address+":"+port, timeout)
	elapsed := time.Since(start)
	if err != nil {
		fmt.Printf("%sConnection timedout...%s\n", Red, Reset)
		return
	}
	defer conn.Close()

	elapsedTime := float64(elapsed.Nanoseconds()) / float64(time.Millisecond)
	elapsedTime = math.Round(elapsedTime*100) / 100

	protocol := fmt.Sprintf("%-5s", "protocol=\x1b[32mTCP\x1b[0m")

	fmt.Printf("Connected to %s%s%s: time=%s%.2fms%+2s %s port=%s%s%s\n", Green, address, Reset, Green, elapsedTime, Reset, protocol, Green, port, Reset)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <address> <port>\n", os.Args[0])
		os.Exit(1)
	}
	address := os.Args[1]
	port := os.Args[2]

	fmt.Printf("paping v1 - By yutho\n\n")

	fmt.Printf("Connecting to  %s%s%s on TCP %s%s%s:\n\n", Green, address, Reset, Green, port, Reset)

	timeout := 1 * time.Second // Adjust timeout as needed

	for {
		papingTCP(address, port, timeout)
		time.Sleep(1 * time.Second)
	}
}
