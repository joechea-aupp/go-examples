package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// c := &http.Client{
	// 	Transport: &retryRoundTripper{
	// 		next: &loggingRoundTripper{
	// 			next:   http.DefaultTransport,
	// 			logger: os.Stdout,
	// 		},
	// 		maxRetries: 3,
	// 		delay:      time.Duration(1 * time.Second),
	// 	},
	// }

	authedClient := &http.Client{
		Transport: &authRoundTripper{
			next: &retryRoundTripper{
				next: &loggingRoundTripper{
					next:   http.DefaultTransport,
					logger: os.Stdout,
				},
				maxRetries: 3,
				delay:      time.Duration(1 * time.Second),
			},
			user: "bob",
			pwd:  "pwd",
		},
	}
	// req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	// req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/status/500", nil)
	req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/basic-auth/bob/pwd", nil)
	if err != nil {
		panic(err)
	}

	res, err := authedClient.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("\n---Response---")
	fmt.Println("STATUS CODE: ", res.StatusCode)
	fmt.Println("BODY: ", string(body))
}

type authRoundTripper struct {
	next      http.RoundTripper
	user, pwd string
}

func (a authRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(a.user, a.pwd)
	return a.next.RoundTrip(r)
}

type loggingRoundTripper struct {
	next   http.RoundTripper
	logger io.Writer
}

// roundtrip is a decorator pattern
func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL.String())
	return l.next.RoundTrip(r)
}

type retryRoundTripper struct {
	next       http.RoundTripper
	maxRetries int
	delay      time.Duration
}

func (rr retryRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	var attempts int
	for {
		res, err := rr.next.RoundTrip(r)
		attempts++

		// max retries exceeded
		if attempts == rr.maxRetries {
			return res, err
		}

		// good outcome
		if err == nil && res.StatusCode < http.StatusInternalServerError {
			return res, err
		}

		// delay and retry
		select {
		// in case request use context and it's timeout
		case <-r.Context().Done():
			return res, r.Context().Err()

		case <-time.After(rr.delay):
		}
	}
}
