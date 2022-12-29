package main

type TraceTempo struct {
	ID            string        `json:"Id"`
	TraceID       string        `json:"traceId"`
	Duration      int64         `json:"duration"`
	Name          string        `json:"name"`
	Timestamp     int64         `json:"timestamp"`
	Tags          Tags          `json:"tags"`
	LocalEndpoint LocalEndpoint `json:"localEndpoint"`
}

type TraceTempoChild struct {
	ID            string        `json:"Id"`
	TraceID       string        `json:"traceId"`
	ParentId      string        `json:"parentId"`
	Duration      int64         `json:"duration"`
	Name          string        `json:"name"`
	Timestamp     int64         `json:"timestamp"`
	LocalEndpoint LocalEndpoint `json:"localEndpoint"`
}

type Tags struct {
	HttpMethod string `json:"http.method"`
	HttpPath   string `json:"http.path"`
}

type LocalEndpoint struct {
	ServiceName string `json:"serviceName"`
}
