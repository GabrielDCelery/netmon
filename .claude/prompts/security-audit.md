# Security Audit

Review the codebase for security issues, focusing on:

## Command Injection
- Check all uses of `os/exec` for unsanitized input
- Verify that no user-controlled values are passed to shell commands
- Ensure command arguments are passed as separate args, not concatenated strings

## Information Exposure
- Check that error messages don't leak sensitive system information
- Verify that process names/PIDs from ss output are handled appropriately
- Review any logging for sensitive data

## Dependency Review
- Run `go list -m all` and check for known vulnerabilities
- Verify all dependencies are from trusted sources
- Check for outdated dependencies with known CVEs

## Report Format
For each finding, provide:
1. File and line number
2. Severity (Critical / High / Medium / Low)
3. Description of the issue
4. Recommended fix
