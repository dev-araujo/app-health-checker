package main

import (
	"fmt"
	"net"
	"time"
)

type HealthCheckResponse struct {
	Status   string `json:"status"`
	Dominio  string `json:"dominio"`
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	ErrorMsg string `json:"error,omitempty"`
}

func Check(destination string, port string) HealthCheckResponse {
	address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp4", address, timeout)
	var response HealthCheckResponse
	if err != nil {
		response = HealthCheckResponse{
			Status:   fmt.Sprintf("ERROR"),
			Dominio:  destination,
			ErrorMsg: fmt.Sprintf("Error: %v", err),
		}
	} else {
		localAddr := conn.LocalAddr().(*net.TCPAddr).IP.String()
		remoteAddr := conn.RemoteAddr().(*net.TCPAddr).IP.String()
		response = HealthCheckResponse{
			Status:  "[SUCCESS]",
			Dominio: destination,
			From:    localAddr,
			To:      remoteAddr,
		}
	}
	return response
}
