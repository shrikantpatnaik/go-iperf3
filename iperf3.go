package iperf3

import (
	"bytes"
	"encoding/json"
	"log"
	"os/exec"
	"strconv"
)

// IperfClient contains all the information required to run IPerf
type IperfClient struct {
	ExecutablePath string
	Host           string
	Port           int
	UDP            bool
	Time           int
	Streams        int
	Reverse        bool
	ForceIPv4      bool
	ForceIPv6      bool
	flags          []string
}

// Init initializes all the flags for running the Iperf command
func (c *IperfClient) Init() {
	c.flags = []string{"-J"}
	if c.UDP {
		c.flags = append(c.flags, "-u")
	}
	if c.Reverse {
		c.flags = append(c.flags, "-r")
	}
	if c.ForceIPv4 {
		c.flags = append(c.flags, "-4")
	} else if c.ForceIPv6 {
		c.flags = append(c.flags, "-6")
	}
	if c.Streams > 0 {
		c.flags = append(c.flags, "-P", strconv.Itoa(c.Streams))
	}
	if c.Time > 0 {
		c.flags = append(c.flags, "-t", strconv.Itoa(c.Time))
	}
	if c.Port > 0 {
		c.flags = append(c.flags, "-p", strconv.Itoa(c.Port))
	}
	c.flags = append(c.flags, "-c", c.Host)
}

// Run runs the iperf command
func (c *IperfClient) Run() ClientResult {
	cmd := exec.Command(c.ExecutablePath, c.flags...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Println("Error: ", err)
	}
	var result ClientResult
	json.Unmarshal(out.Bytes(), &result)
	return result
}

// NewIperf3Client creates a new iperf3 client with the given host
func NewIperf3Client(host string) IperfClient {
	return IperfClient{
		ExecutablePath: "iperf3",
		Host:           host,
		Port:           0,
		UDP:            false,
		Time:           0,
		Streams:        0,
		Reverse:        false,
		ForceIPv4:      false,
		ForceIPv6:      false,
	}
}
