package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/simonski/cli"
)

const DEFAULT_MIN = 1024
const DEFAULT_MAX = 65000
const DEFAULT_STEP = 10

func main() {
	cli := cli.New(os.Args)

	min := cli.GetIntOrEnvOrDefault("-min", "RPORT_MIN", DEFAULT_MIN)
	max := cli.GetIntOrEnvOrDefault("-max", "RPORT_MAX", DEFAULT_MAX)
	step := cli.GetIntOrEnvOrDefault("-step", "RPORT_STEP", DEFAULT_STEP)

	for port := min; port <= max; port += step {
		portstr := fmt.Sprintf("%v", port)
		if IsTCPPortAvailable("localhost", portstr) {
			fmt.Println(portstr)
			break
		}
	}
	os.Exit(1)

}

func IsTCPPortAvailable(host string, port string) bool {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", host+":"+port, timeout)
	if err != nil {
		return true
	}
	conn.Close()
	return false
}
