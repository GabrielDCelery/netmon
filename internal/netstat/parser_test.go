package netstat

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Connection
	}{
		{
			name:     "empty input",
			input:    "",
			expected: nil,
		},
		{
			name:     "header only",
			input:    "Netid  State      Recv-Q Send-Q Local Address:Port   Peer Address:Port  Process",
			expected: nil,
		},
		{
			name: "single TCP ESTAB connection",
			input: `Netid  State      Recv-Q Send-Q Local Address:Port   Peer Address:Port  Process
tcp    ESTAB      0      0      192.168.1.100:42356  93.184.216.34:443   users:(("firefox",pid=1234,fd=45))`,
			expected: []Connection{
				{
					Protocol: "tcp",
					State:    "ESTAB",
					RecvQ:    "0",
					SendQ:    "0",
					Local:    "192.168.1.100:42356",
					Peer:     "93.184.216.34:443",
					Process:  `users:(("firefox",pid=1234,fd=45))`,
				},
			},
		},
		{
			name: "multiple connections",
			input: `Netid  State      Recv-Q Send-Q Local Address:Port   Peer Address:Port  Process
tcp    ESTAB      0      0      192.168.1.100:42356  93.184.216.34:443   users:(("firefox",pid=1234,fd=45))
tcp    LISTEN     0      128    0.0.0.0:22            0.0.0.0:*           users:(("sshd",pid=789,fd=3))
udp    UNCONN     0      0      127.0.0.53%lo:53      0.0.0.0:*           users:(("systemd-resolve",pid=567,fd=12))`,
			expected: []Connection{
				{
					Protocol: "tcp",
					State:    "ESTAB",
					RecvQ:    "0",
					SendQ:    "0",
					Local:    "192.168.1.100:42356",
					Peer:     "93.184.216.34:443",
					Process:  `users:(("firefox",pid=1234,fd=45))`,
				},
				{
					Protocol: "tcp",
					State:    "LISTEN",
					RecvQ:    "0",
					SendQ:    "128",
					Local:    "0.0.0.0:22",
					Peer:     "0.0.0.0:*",
					Process:  `users:(("sshd",pid=789,fd=3))`,
				},
				{
					Protocol: "udp",
					State:    "UNCONN",
					RecvQ:    "0",
					SendQ:    "0",
					Local:    "127.0.0.53%lo:53",
					Peer:     "0.0.0.0:*",
					Process:  `users:(("systemd-resolve",pid=567,fd=12))`,
				},
			},
		},
		{
			name: "connection without process info",
			input: `Netid  State      Recv-Q Send-Q Local Address:Port   Peer Address:Port  Process
tcp    TIME-WAIT  0      0      192.168.1.100:54321  10.0.0.1:80`,
			expected: []Connection{
				{
					Protocol: "tcp",
					State:    "TIME-WAIT",
					RecvQ:    "0",
					SendQ:    "0",
					Local:    "192.168.1.100:54321",
					Peer:     "10.0.0.1:80",
					Process:  "",
				},
			},
		},
		{
			name: "IPv6 connections",
			input: `Netid  State      Recv-Q Send-Q Local Address:Port   Peer Address:Port  Process
tcp    LISTEN     0      128    [::]:80               [::]:*              users:(("nginx",pid=1000,fd=6))
tcp    ESTAB      0      0      [::1]:5432            [::1]:43210         users:(("postgres",pid=2000,fd=8))`,
			expected: []Connection{
				{
					Protocol: "tcp",
					State:    "LISTEN",
					RecvQ:    "0",
					SendQ:    "128",
					Local:    "[::]:80",
					Peer:     "[::]:*",
					Process:  `users:(("nginx",pid=1000,fd=6))`,
				},
				{
					Protocol: "tcp",
					State:    "ESTAB",
					RecvQ:    "0",
					SendQ:    "0",
					Local:    "[::1]:5432",
					Peer:     "[::1]:43210",
					Process:  `users:(("postgres",pid=2000,fd=8))`,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Parse(tt.input)

			if tt.expected == nil {
				if result != nil {
					t.Errorf("expected nil, got %v", result)
				}
				return
			}

			if len(result) != len(tt.expected) {
				t.Fatalf("expected %d connections, got %d", len(tt.expected), len(result))
			}

			for i, exp := range tt.expected {
				got := result[i]
				if got.Protocol != exp.Protocol {
					t.Errorf("connection[%d].Protocol = %q, want %q", i, got.Protocol, exp.Protocol)
				}
				if got.State != exp.State {
					t.Errorf("connection[%d].State = %q, want %q", i, got.State, exp.State)
				}
				if got.RecvQ != exp.RecvQ {
					t.Errorf("connection[%d].RecvQ = %q, want %q", i, got.RecvQ, exp.RecvQ)
				}
				if got.SendQ != exp.SendQ {
					t.Errorf("connection[%d].SendQ = %q, want %q", i, got.SendQ, exp.SendQ)
				}
				if got.Local != exp.Local {
					t.Errorf("connection[%d].Local = %q, want %q", i, got.Local, exp.Local)
				}
				if got.Peer != exp.Peer {
					t.Errorf("connection[%d].Peer = %q, want %q", i, got.Peer, exp.Peer)
				}
				if got.Process != exp.Process {
					t.Errorf("connection[%d].Process = %q, want %q", i, got.Process, exp.Process)
				}
			}
		})
	}
}
