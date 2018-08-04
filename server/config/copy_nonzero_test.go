package config

import (
	"testing"
)

var server1 = Server{
	Port:    123,
	IP:      "127.0.0.1",
	LogFile: "/var/log/chat-server.log",
}

func TestCopyNonzeroCopyAll(t *testing.T) {
	server2 := Server{}

	err := copyNonzero(server1, &server2)
	if err != nil {
		t.Fatalf("Error in copyNonzero: %v", err)
	}

	if server2.Port != server1.Port {
		t.Errorf("Port not copied, it is '%v', should be '%v'", server2.Port, server1.Port)
	}

	if server2.IP != server1.IP {
		t.Errorf("IP not copied, it is '%v', should be '%v'", server2.IP, server1.IP)
	}

	if server2.LogFile != server1.LogFile {
		t.Errorf("LogFile not copied, it is '%v', should be '%v'", server2.LogFile, server1.LogFile)
	}
}

func TestCopyNonzeroCopyNone(t *testing.T) {
	port := 246
	ip := "127.0.0.254"
	logFile := "~/.chat.log"

	server2 := Server{
		Port:    port,
		IP:      ip,
		LogFile: logFile,
	}

	err := copyNonzero(server1, &server2)
	if err != nil {
		t.Fatalf("Error in copyNonzero: %v", err)
	}

	if server2.Port != port {
		t.Errorf("Port not copied, it is '%v', should be '%v'", server2.Port, port)
	}

	if server2.IP != ip {
		t.Errorf("IP not copied, it is '%v', should be '%v'", server2.IP, ip)
	}

	if server2.LogFile != logFile {
		t.Errorf("LogFile not copied, it is '%v', should be '%v'", server2.LogFile, logFile)
	}
}

func TestCopyNonzeroCopySome(t *testing.T) {
	logFile := "~/.chat.log"

	server2 := Server{
		LogFile: logFile,
	}

	err := copyNonzero(server1, &server2)
	if err != nil {
		t.Fatalf("Error in copyNonzero: %v", err)
	}

	if server2.Port != server1.Port {
		t.Errorf("Port not copied, it is '%v', should be '%v'", server2.Port, server1.Port)
	}

	if server2.IP != server1.IP {
		t.Errorf("IP not copied, it is '%v', should be '%v'", server2.IP, server1.IP)
	}

	if server2.LogFile != logFile {
		t.Errorf("LogFile not copied, it is '%v', should be '%v'", server2.LogFile, logFile)
	}
}

func TestCopyNonzeroDeepCopy(t *testing.T) {
	logFile := "~/.chat.log"

	server2 := Server{
		LogFile: logFile,
	}

	type TestStruct struct {
		S Server
		I int
		f float64
	}

	s1 := TestStruct{
		S: server1,
		I: 123,
		f: 3.14,
	}

	s2 := TestStruct{
		S: server2,
		I: 0,
		f: 1.0,
	}

	err := copyNonzero(s1, &s2)
	if err != nil {
		t.Fatalf("Error in copyNonzero: %v", err)
	}

	if s2.S.Port != s1.S.Port {
		t.Errorf("Port not copied, it is '%v', should be '%v'", s2.S.Port, s1.S.Port)
	}

	if s2.S.IP != s1.S.IP {
		t.Errorf("IP not copied, it is '%v', should be '%v'", s2.S.IP, s1.S.IP)
	}

	if s2.S.LogFile != logFile {
		t.Errorf("LogFile not copied, it is '%v', should be '%v'", s2.S.LogFile, logFile)
	}

	if s2.I != s1.I {
		t.Errorf("i not copied, it is '%v', should be '%v'", s2.I, s1.I)
	}
}
