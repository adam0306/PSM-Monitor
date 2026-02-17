package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://<psm>/psm/api/health")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}


package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// ----- Configuration ----------------------------------------------------
	// Change these values as needed
	url := "https://example.com" // target address
	keyword := "Go"              // filter: keep lines containing this word
	// -----------------------------------------------------------------------

	// 1️⃣ Perform the GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Check for non‑200 status codes
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "bad status: %s\n", resp.Status)
		os.Exit(1)
	}

	// 2️⃣ Read and filter the response line‑by‑line
	fmt.Printf("Lines from %s containing \"%s\":\n\n", url, keyword)
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadString('\n')
		if strings.Contains(line, keyword) {
			fmt.Print(line) // line matches the filter, print it
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
			os.Exit(1)
		}
	}
}



package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// checkURL contacts the given address, looks for the words “pass” or “fail”
// in the response body, and returns an appropriate error.
// It also propagates network‑level errors (timeouts, DNS failures, etc.).
func checkURL(address string) error {
	// Create a client with a reasonable timeout.
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Build the request – you can customise method, headers, etc. here.
	req, err := http.NewRequest(http.MethodGet, address, nil)
	if err != nil {
		return fmt.Errorf("building request: %w", err)
	}
	// Example header, adjust as needed:
	req.Header.Set("Accept", "text/plain")

	// Perform the request.
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Non‑2xx status codes are usually worth treating as errors.
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected HTTP status: %d %s", resp.StatusCode, resp.Status)
	}

	// Read the body (limit size to avoid OOM on huge responses).
	const maxBody = 1 << 20 // 1 MiB
	body, err := io.ReadAll(io.LimitReader(resp.Body, maxBody))
	if err != nil {
		return fmt.Errorf("reading response: %w", err)
	}

	// Normalise whitespace and case for easier matching.
	clean := strings.TrimSpace(strings.ToLower(string(body)))

	switch {
	case strings.Contains(clean, "pass"):
		// Success – no error to return.
		return nil
	case strings.Contains(clean, "fail"):
		return errors.New("response indicates failure")
	default:
		return errors.New("neither 'pass' nor 'fail' found in response")
	}
}

// ---------------------------------------------------------------------------
// Example usage
func main() {
	url := "https://example.com/healthcheck"

	if err := checkURL(url); err != nil {
		// Handle the error as you see fit – log, retry, abort, etc.
		fmt.Printf("Check failed: %v\n", err)
	} else {
		fmt.Println("Check succeeded!")
	}
}