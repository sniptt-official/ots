/*
Copyright © 2021 Sniptt <support@sniptt.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package client

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/sniptt-official/ots/build"
	"github.com/spf13/viper"
)

// APIError represents an error response from the API
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error: %d - %s", e.StatusCode, e.Message)
}

type CreateOtsReq struct {
	EncryptedBytes string `json:"encryptedBytes"`
	ExpiresIn      uint32 `json:"expiresIn"`
}

type CreateOtsRes struct {
	Id        string `json:"id"`
	ExpiresAt int64  `json:"expiresAt"`
	ViewURL   *url.URL
}

// CreateOts creates a new one-time secret with the given parameters
func CreateOts(ctx context.Context, encryptedBytes []byte, expiresIn time.Duration, region string) (*CreateOtsRes, error) {
	// Get region-specific API URL
	apiURL := viper.GetString(fmt.Sprintf("apiUrl.%s", region))
	if apiURL == "" {
		return nil, fmt.Errorf("no API URL configured for region: %s", region)
	}

	apiKey := viper.GetString("apiKey")

	// Build the request
	reqURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, fmt.Errorf("invalid API URL: %w", err)
	}

	reqBody := &CreateOtsReq{
		EncryptedBytes: base64.StdEncoding.EncodeToString(encryptedBytes),
		ExpiresIn:      uint32(expiresIn.Seconds()),
	}

	// Using json.Marshal instead of Encoder since we have the full payload
	payload, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL.String(), bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Client-Name", "ots-cli")
	req.Header.Set("X-Client-Version", build.Version)

	// Add optional authentication (for self-hosted)
	if apiKey != "" {
		req.Header.Set("X-Api-Key", apiKey)
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer res.Body.Close()

	// Build the response
	resBody := &CreateOtsRes{}

	if err := decodeJSON(res, resBody); err != nil {
		return nil, err
	}

	viewURL, err := url.Parse(res.Header.Get("X-View-Url"))
	if err != nil {
		return nil, fmt.Errorf("invalid view URL: %w", err)
	}

	resBody.ViewURL = viewURL

	return resBody, nil
}

func decodeJSON(res *http.Response, target interface{}) error {
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("failed to read error response: %w", err)
		}

		return &APIError{
			StatusCode: res.StatusCode,
			Message:    string(body),
		}
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
