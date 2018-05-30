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
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setAliasCmd represents the setAlias command
var setAliasCmd = &cobra.Command{
	Use:   "set-alias",
	Short: "You can name tx hash.",
	Long: `With this command you can give a name to the transaction hash in which the Shell script is written.

For example:
nem.sh set-alias --hash b37685ca16474b6897550f51f008c11b1e24e93e51b5543d066d9266d4e35008 --name hello
`,
	Run: func(cmd *cobra.Command, args []string) {
		if aHash != "" && name != "" {
			tmp := viper.Get("alias")
			h := tmp.(map[string]interface{})
			if h == nil {
				h = make(map[string]interface{})
			}
			h[name] = aHash
			viper.Set("alias", h)
			home, _ := homedir.Dir()
			err := viper.WriteConfigAs(home + "/.nem.sh/config.json")
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}

var aHash string
var name string

func init() {
	rootCmd.AddCommand(setAliasCmd)
	setAliasCmd.Flags().StringVarP(&aHash, "hash", "", "", "Source directory to read from")
	setAliasCmd.Flags().StringVarP(&name, "name", "", "", "Source directory to read from")
}
