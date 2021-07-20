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
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type CreateOtsReq struct {
	EncryptedBytes string `json:"EncryptedBytes"`
	ExpiresIn      uint32 `json:"ExpiresIn"`
}

type CreateOtsRes struct {
	ID        string `json:"Id"`
	ExpiresAt int64  `json:"ExpiresAt"`
	ViewURL   *url.URL
}

func CreateOts(encryptedBytes []byte, expiresIn uint32) (*CreateOtsRes, error) {
	reqBody := &CreateOtsReq{
		EncryptedBytes: base64.StdEncoding.EncodeToString(encryptedBytes),
		ExpiresIn:      expiresIn, // Seconds resolution
	}

	resBody := &CreateOtsRes{}

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(reqBody)
	if err != nil {
		return nil, err
	}

	// TODO: Make part of config
	res, err := http.Post("https://apiv2.beta.snipt.io/secrets", "application/json", payloadBuf)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = decodeJSON(res, &resBody)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(res.Header.Get("X-View-Url"))
	if err != nil {
		return nil, err
	}

	resBody.ViewURL = u

	return resBody, nil
}

func decodeJSON(res *http.Response, target interface{}) error {
	statusOK := res.StatusCode >= 200 && res.StatusCode < 300

	if !statusOK {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		// TODO: Implement error struct response.
		return errors.New(string(body))
	}

	return json.NewDecoder(res.Body).Decode(target)
}
