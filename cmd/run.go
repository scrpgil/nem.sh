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
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/myndshft/nemgo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"os/exec"
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
			tmp := viper.Get("alias")
			if tmp != nil {
				a := tmp.(map[string]interface{})
				tmp2 := a[alias]
				if tmp2 != nil {
					h = a[alias].(string)
				} else {
					fmt.Println("alias not found.")
					return
				}
			} else {
				fmt.Println("alias not found.")
				return
			}
		} else {
			h = hash
		}
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
			if view == true {
				fmt.Println(string(decoded))
			} else {
				out, err := exec.Command("sh", "-c", string(decoded)).Output()
				if err != nil {
					fmt.Println(err)
				}
				fmt.Printf("%s", out)
			}
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
