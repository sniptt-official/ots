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
package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"regexp"
	"testing"
)

func TestBytes(t *testing.T) {
	secret := "nuclear_launch_codes"
	want := regexp.MustCompile(secret)

	bytesToEncrypt := []byte(secret)
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	nonce, _ := hex.DecodeString("64a9433eae7ccceee2fc0eda")

	// Encrypt secret.
	ciphertext, secretKey, err := Bytes(BytesParams{Bytes: bytesToEncrypt, EncryptionKey: key, Nonce: nonce})

	// Remove nonce from start of ciphertext.
	ciphertext = ciphertext[len(nonce):]

	if bytes.Compare(key, secretKey) != 0 || err != nil {
		t.Fatalf("Expected %v, got %v", key, secretKey)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	decryptedSecret, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	if !want.MatchString(string(decryptedSecret)) {
		t.Fatalf("Expected %v, got %v", want, decryptedSecret)
	}
}
