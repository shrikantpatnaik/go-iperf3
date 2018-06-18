package iperf3

// ClientResult contains results from the iperf run
type ClientResult struct {
	Start     Start       `json:"start"`
	Intervals []Intervals `json:"intervals"`
	End       End         `json:"end"`
}

// Start contains all the Start information from the iperf run
type Start struct {
	Connection    []StartConnected  `json:"connected"`
	Version       string            `json:"version"`
	SystemInfo    string            `json:"system_info"`
	Timestamp     StartTimestamp    `json:"timestamp"`
	ConnectingTo  StartConnectingTo `json:"connecting_to"`
	Cookie        string            `json:"cookie"`
	TCPMSSDefault int               `json:"tcp_mss_default"`
	SockBufSize   int               `json:"sock_bufsize"`
	SndBufActual  int               `json:"sndbuf_actual"`
	RcvBufActual  int               `json:"rcvbuf_actual"`
	TestStart     StartTestStart    `json:"test_start"`
}

// StartConnected contains all the connected information from the Start of the iperf run
type StartConnected struct {
	Socket     int    `json:"socket"`
	LocalHost  string `json:"local_host"`
	LocalPort  int    `json:"local_port"`
	RemoteHost string `json:"remote_host"`
	RemotePort int    `json:"remote_port"`
}

// StartTimestamp contains all the timestamp information from the Start of the iperf run
type StartTimestamp struct {
	Time     string `json:"time"`
	TimeSecs int    `json:"timesecs"`
}

// StartConnectingTo contains all the connecting_to information from the Start of the iperf run
type StartConnectingTo struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// StartTestStart contains all the test_start information from the Start of the iperf run
type StartTestStart struct {
	Protocol string `json:"protocol"`
	Streams  int    `json:"num_streams"`
	BlkSize  int    `json:"blksize"`
	Omit     int    `json:"omit"`
	Duration int    `json:"duration"`
	Bytes    int    `json:"bytes"`
	Blocks   int    `json:"blocks"`
	Reverse  int    `json:"reverse"`
	Tos      int    `json:"tos"`
}

// Intervals contains all the intervals information from the iperf run
type Intervals struct {
	Streams []Stream `json:"streams"`
	Sum     Sum      `json:"sum"`
}

// End contains all the end information from the iperf run
type End struct {
	Streams        []EndStreams          `json:"streams"`
	SumSent        Sum                   `json:"sum_sent"`
	SumReceived    Sum                   `json:"sum_received"`
	CPUUtilization CPUUtilizationPercent `json:"cpu_utilization_percent"`
}

// EndStreams contains all the stream information from the end of the iperf run
type EndStreams struct {
	Sender   Stream `json:"sender"`
	Receiver Stream `json:"receiver"`
}

// CPUUtilizationPercent contains the CPU Utilization information from the end of the iperf run
type CPUUtilizationPercent struct {
	HostTotal    float64 `json:"host_total"`
	HostUser     float64 `json:"host_user"`
	HostSystem   float64 `json:"host_system"`
	RemoteTotal  float64 `json:"remote_total"`
	RemoteUser   float64 `json:"remote_user"`
	RemoteSystem float64 `json:"remote_system"`
}

// Stream contains all the information for each stream
type Stream struct {
	Socket        int     `json:"socket"`
	Start         float64 `json:"start"`
	End           float64 `json:"end"`
	Seconds       float64 `json:"seconds"`
	Bytes         float64 `json:"bytes"`
	BitsPerSecond float64 `json:"bits_per_second"`
	Omitted       bool    `json:"omitted"`
}

// Sum contains all the Sum information for each stream
type Sum struct {
	Start         float64 `json:"start"`
	End           float64 `json:"end"`
	Seconds       float64 `json:"seconds"`
	Bytes         float64 `json:"bytes"`
	BitsPerSecond float64 `json:"bits_per_second"`
	Omitted       bool    `json:"omitted"`
}
