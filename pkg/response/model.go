package response

type Message struct {
	ErrorLog  ErrorLogs `json:"errorLog,omitempty"`
	HostName  string    `json:"hostName,omitempty"`
	Status    string    `json:"status,omitempty"`
	TimeTaken string    `json:"timeTaken,omitempty"`
	Count     int       `json:"count,omitempty"`
}

type ErrorLogs []ErrorLog

type ErrorLog struct {
	Status    string `json:"status,omitempty"`
	Trace     string `json:"trace,omitempty"`
	RootCause string `json:"rootCause,omitempty"`
}
