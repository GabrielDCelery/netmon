# Task Backlog

## Task 1: ss/netstat TUI View

Build the primary TUI view for displaying network connections from `ss -tunap`.

### Requirements

- Table view with columns: Protocol, State, Recv-Q, Send-Q, Local Address, Peer Address, Process
- Color-coded connection states (ESTAB=green, LISTEN=blue, TIME-WAIT=yellow, CLOSE-WAIT=red)
- Auto-refresh every 2 seconds
- Keyboard navigation (up/down arrows, page up/down)
- Filtering by state, protocol, or address (planned)
- Sorting by column (planned)
- Status bar showing connection count and last refresh time

### Acceptance Criteria

- [ ] Table renders correctly with real ss output
- [ ] States are color-coded
- [ ] Auto-refresh works without flickering
- [ ] Arrow keys navigate the table
- [ ] q/ctrl+c quits cleanly
- [ ] Parser tests cover all connection states
- [ ] Works on Linux (graceful error on other platforms)
