package netstat

import (
	"strings"
)

// Parse takes raw ss -tunap output and returns a slice of Connection structs.
func Parse(raw string) []Connection {
	lines := strings.Split(strings.TrimSpace(raw), "\n")
	if len(lines) < 2 {
		return nil
	}

	var connections []Connection

	// Skip header line
	for _, line := range lines[1:] {
		conn, ok := parseLine(line)
		if ok {
			connections = append(connections, conn)
		}
	}

	return connections
}

func parseLine(line string) (Connection, bool) {
	fields := strings.Fields(line)
	if len(fields) < 6 {
		return Connection{}, false
	}

	conn := Connection{
		Protocol: fields[0],
		State:    fields[1],
		RecvQ:    fields[2],
		SendQ:    fields[3],
		Local:    fields[4],
		Peer:     fields[5],
	}

	if len(fields) > 6 {
		conn.Process = strings.Join(fields[6:], " ")
	}

	return conn, true
}
