# go-iperf3

A simple wrapper to run the iperf3 executable as a client and convert all returned information to Go Structs so that the information is easily accessible.

## Requirements

iperf3 must be installed on the system, if it is not available in the PATH, then it must be explicitly set using client.ExecutablePath. See the example below for more.

## Installation

`go get -u github.com/shrikantpatnaik/go-iperf3`

## Usage

### Simple

```go
package main

import (
	"log"

	iperf3 "github.com/shrikantpatnaik/go-iperf3"
)

func main() {
	// Create a client
	client := iperf3.NewIperf3Client("bouygues.iperf.fr")

	// Initialize the client, if any options are being set then this must be called after
	client.Init()

	//Run the client
	result, err := client.Run()
	if err != nil {
		log.Println("Error: ", err)
		return
	}

	log.Printf("%+v\n", result)
}
```

### All Options

All the options default to the default values in the executable by not explicitly specifying them.

```go
package main

import (
	"log"

	iperf3 "github.com/shrikantpatnaik/go-iperf3"
)

func main() {
	client := iperf3.NewIperf3Client("ping.online.net")

	// Set the path to the iperf3 executable
	client.ExecutablePath = "/usr/local/bin/iperf3"

	// Set the port for the iperf3 server, defaults to 5201
	client.Port = 55201

	// Use UDP instead of TCP, defaults to false
	client.UDP = true // Defaults to false

	// Set the time in seconds to transmit for, defaults to 10 seconds
	client.Time = 1 // Defaults to 10

	// Set the number of parallel client streams to use, defaults to 1
	client.Streams = 10

	// Run in reverse mode (server sends, client receives), defaults to false
	client.Reverse = true

	// Force IPv4, defaults to false, executable automatically determines whether to use ipv4 or ipv6
	client.ForceIPv4 = true

	// Force IPv6, defaults to false, executable automatically determines whether to use ipv4 or ipv6
	client.ForceIPv6 = true // This is ignored if ForceIPv4 is set to true

	client.Init()
	result, err := client.Run()
	if err != nil {
		log.Println("Error: ", err)
		return
	}

	log.Printf("%+v\n", result)
}
```
