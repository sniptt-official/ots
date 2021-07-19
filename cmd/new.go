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
package cmd

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"time"

	"github.com/sniptt-official/ots/api/client"
	"github.com/sniptt-official/ots/crypto/encrypt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var expires string

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create end-to-end encrypted secret",
	Long: `
Encrypts a secret and makes it available for sharing via one-time URL.

The secret is stored encrypted for a specified duration which can range
from 5 minutes to 7 days (default is 72 hours). The secret gets deleted
from the server upon retrieval therefore can only be viewed once.
`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		duration, err := time.ParseDuration(expires)
		if err != nil {
			return err
		}

		if duration.Minutes() < 5 {
			return errors.New("expiry must be at least 5 minutes")
		}

		if duration.Hours() > 168 {
			return errors.New("expiry must be less than 7 days")
		}

		bytes, err := getInputBytes()
		if err != nil {
			return err
		}

		encryptedBytes, secretKey, err := encrypt.Bytes(bytes)
		if err != nil {
			return err
		}

		ots, err := client.CreateOts(encryptedBytes, uint32(duration.Seconds()))
		if err != nil {
			return err
		}

		expiresAt := time.Unix(ots.ExpiresAt, 0)

		q := ots.ViewUrl.Query()
		q.Set("p", base64.URLEncoding.EncodeToString(secretKey))
		q.Set("ref", "cli")
		ots.ViewUrl.RawQuery = q.Encode()

		msg := fmt.Sprintf(`
Your secret is now available on the below URL.

%v

You should only share this URL with the intended recipient.

Please note that once retrieved, the secret will no longer
be available for viewing. If not viewed, the secret will
automatically expire at approximately %v.
`,
			ots.ViewUrl,
			expiresAt.Format("2 Jan 2006 15:04:05"),
		)

		fmt.Print(msg)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&expires, "expires", "x", "24h", "Secret will be deleted from the server after specified duration, supported units: s,m,h")
}

func getInputBytes() ([]byte, error) {
	fi, _ := os.Stdin.Stat() // get the FileInfo struct describing the standard input.

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		bytes, err := ioutil.ReadAll(os.Stdin)

		if err != nil {
			return nil, err
		}

		return bytes, nil
	} else {
		fmt.Print("Enter your secret: ")

		bytes, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return nil, err
		}

		return []byte(bytes), nil
	}
}
