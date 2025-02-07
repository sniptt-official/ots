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
package encrypt_test

import (
	"crypto/aes"
	"crypto/cipher"
	"testing"

	"github.com/sniptt-official/ots/crypto/encrypt"
)

func TestBytes(t *testing.T) {
	t.Run("successfully encrypts and decrypts data", func(t *testing.T) {
		unencryptedSecret := "nuclear_launch_codes"

		encryptedBytes, err := encrypt.Bytes([]byte(unencryptedSecret))
		if err != nil {
			t.Fatalf("encryption failed: %v", err)
		}

		// Use helper methods to get ciphertext without nonce
		ciphertext := encryptedBytes.CiphertextWithoutNonce()
		nonce := encryptedBytes.ExtractNonce()

		block, err := aes.NewCipher(encryptedBytes.Key)
		if err != nil {
			t.Fatalf("failed to create cipher: %v", err)
		}

		aesGCM, err := cipher.NewGCM(block)
		if err != nil {
			t.Fatalf("failed to create GCM: %v", err)
		}

		decryptedSecret, err := aesGCM.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			t.Fatalf("decryption failed: %v", err)
		}

		if got := string(decryptedSecret); got != unencryptedSecret {
			t.Errorf("got %q, want %q", got, unencryptedSecret)
		}
	})

	t.Run("returns error for empty input", func(t *testing.T) {
		_, err := encrypt.Bytes([]byte{})
		if err == nil {
			t.Error("expected error for empty input, got nil")
		}
	})
}
