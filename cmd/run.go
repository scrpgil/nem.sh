// Copyright © 2018 xiaca@scrpgil <scraplbuild@gmai.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/myndshft/nemgo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

var hash string
var alias string
var view bool

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute the shell script written in the transaction message.",
	Long: `This command executes the shell script written in the transaction message.
	
	example:
	nem.sh run --hash b37685ca16474b6897550f51f008c11b1e24e93e51b5543d066d9266d4e35008
	 or
	nem.sh run --name hello
	 or
	nem.sh run --name hello --view
	`,
	Run: func(cmd *cobra.Command, args []string) {
		address := viper.GetString("address")
		h := ""
		if alias != "" {
			if tmp, err := getAlias(alias); err != nil {
				fmt.Println(err)
				return
			} else {
				h = tmp
			}
		} else {
			h = hash
		}

		// GetFilePath
		filePath := GetFilePath(h)
		if Exists(filePath) == false {
			byteArray, _ := request(h)
			t := &nemgo.TransactionMetadataPair{}
			if err := json.Unmarshal(byteArray, t); err != nil {
				fmt.Println("JSON Unmarshal error:", err)
				return
			}
			if t.Meta.Hash.Data == h && t.Transaction.Recipient == address {
				p := t.Transaction.Message.Payload
				decoded, err := hex.DecodeString(p)
				if err != nil {
					fmt.Println(err)
				}
				var buf bytes.Buffer
				buf.Write(decoded)
				if err := ioutil.WriteFile(filePath, buf.Bytes(), 0644); err != nil {
					fmt.Println("err:", err)
				}
			}
		}

		if view == true {
			out, err := exec.Command("cat", filePath).Output()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%s\n", out)
		} else {
			str := strings.Join(args, " ")
			out, err := exec.Command("sh", filePath, str).Output()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%s", out)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&hash, "hash", "", "", "Source directory to read from")
	runCmd.Flags().StringVarP(&alias, "alias", "", "", "Source directory to read from")
	runCmd.Flags().BoolVar(&view, "view", false, "Source directory to read from")
}

func request(hash string) ([]byte, error) {
	node := viper.GetString("node")
	port := viper.GetString("port")
	p := viper.GetString("protocol")
	url := p + "://" + node + ":" + port + "/transaction/get?hash=" + hash
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	return byteArray, err
}

func getAlias(key string) (string, error) {
	tmp := viper.Get("alias")
	if tmp != nil {
		a := tmp.(map[string]interface{})
		tmp2 := a[key]
		if tmp2 != nil {
			h := a[key].(string)
			return h, nil
		} else {
			return "", errors.New("alias not found.")
		}
	} else {
		return "", errors.New("alias not found.")
	}
}
