package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

const maxRetries = 20
const waitBetweenRetries = time.Second * 1

func HttpGetWithRetry[T any](name, url string) (T, error) {
	var result T
	var response *http.Response
	var err error
	for i := 0; i < maxRetries; i++ {
		client := &http.Client{Timeout: time.Second * 2}
		response, err = client.Get(url)
		if err != nil {
			// EOF is a common error when the server is not ready to accept requests yet
			if errors.Is(err, io.EOF) || isTimeout(err) {
				time.Sleep(waitBetweenRetries)
				continue
			}
			return result, fmt.Errorf("unable to query %s API: %w", name, err)
		}
		defer response.Body.Close()

		// Accepted means that the API has registered the request and we need to check again later
		// for the response
		if response.StatusCode == http.StatusAccepted {
			time.Sleep(waitBetweenRetries)
			continue
		}
		if response.StatusCode == http.StatusOK {
			break
		}
	}

	if response.StatusCode != http.StatusOK {
		return result, fmt.Errorf("unable to query forecast API after retries status code: %s", response.Status)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return result, fmt.Errorf("unable to read body from %s API: %w", name, err)
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, fmt.Errorf("unable to decode JSON from %s API: %w", name, err)
	}
	return result, nil
}

func isTimeout(err error) bool {
	nerr, ok := err.(net.Error)
	return ok && nerr.Timeout()
}
