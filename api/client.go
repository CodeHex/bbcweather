package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const maxRetries = 20
const waitBetweenRetries = time.Second * 2

func HttpGetWithRetry[T any](name, url string) (T, error) {
	var result T
	var response *http.Response
	var err error
	for i := 0; i < maxRetries; i++ {
		response, err = http.Get(url)
		if err != nil {
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
