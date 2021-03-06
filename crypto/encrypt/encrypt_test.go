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
	"strings"
	"testing"

	"github.com/sniptt-official/ots/crypto/encrypt"
)

func TestBytes(t *testing.T) {
	unencryptedSecret := "nuclear_launch_codes"

	encryptedBytes, err := encrypt.Bytes([]byte(unencryptedSecret))
	if err != nil {
		t.Fatalf(err.Error())
	}

	ciphertext, key, nonce := encryptedBytes.Ciphertext, encryptedBytes.Key, encryptedBytes.Nonce
	ciphertext = ciphertext[len(nonce):] // Remove nonce from start of ciphertext.

	block, err := aes.NewCipher(key)
	if err != nil {
		t.Fatalf(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		t.Fatalf(err.Error())
	}

	decryptedSecret, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if strings.Compare(unencryptedSecret, string(decryptedSecret)) != 0 {
		t.Fatalf(`Bytes(nil, %v) = %v; want: %v`, []byte(unencryptedSecret), string(decryptedSecret), unencryptedSecret)
	}
}
