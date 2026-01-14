package cmdutil

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// outputExec runs the given command and captures its output, decoding from GBK if necessary.
func outputExec(cmd *exec.Cmd) ([]string, error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	// Read stdout
	scanner := bufio.NewScanner(stdout)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, tryDecodeGBKString(line))
	}

	if err := scanner.Err(); err != nil {
		cmd.Wait() // Clean up process on scanner error
		return nil, err
	}

	// Read stderr
	stderrScanner := bufio.NewScanner(stderr)
	var stderrLines []string
	for stderrScanner.Scan() {
		stderrLines = append(stderrLines, tryDecodeGBKString(stderrScanner.Text()))
	}

	if err := stderrScanner.Err(); err != nil {
		cmd.Wait() // Clean up process on scanner error
		return nil, err
	}

	// Wait for command completion and let it handle pipe cleanup
	waitErr := cmd.Wait()

	if waitErr != nil {
		if len(stderrLines) > 0 {
			return lines, fmt.Errorf("%w: stderr: %s", waitErr, strings.Join(stderrLines, "\n"))
		}
		return lines, waitErr
	}

	return lines, nil
}

// tryDecodeGBKString helper function to decode potential GBK encoded strings
func tryDecodeGBKString(input string) string {
	// Only attempt GBK decoding if we detect non-UTF8 byte sequences
	// This prevents unnecessary decoding of already-correct UTF-8 strings

	// First check if the string appears to be valid UTF-8
	if utf8.ValidString(input) {
		// If it's valid UTF-8, it might still contain Chinese characters
		// but no need to decode unless it looks like GBK-encoded bytes in string form
		return input
	}

	// If not valid UTF-8, try to decode as GBK
	decoded, err := simplifiedchinese.GBK.NewDecoder().String(input)
	if err != nil {
		// If decoding fails, return original string
		return input
	}

	// Return decoded version only if successful
	return decoded
}
