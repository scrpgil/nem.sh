// Copyright © 2018 xiaca@scrpgil <scraplbuild@gmai.com>
///
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
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "nem.sh",
	Short: "nem.sh is a tool that executes the shell script written in the transaction message.",
	Long: `
                                                                                      hhhhhhh             
                                                                                      h:::::h             
                                                                                      h:::::h             
                                                                                      h:::::h             
nnnn  nnnnnnnn        eeeeeeeeeeee       mmmmmmm    mmmmmmm               ssssssssss   h::::h hhhhh       
n:::nn::::::::nn    ee::::::::::::ee   mm:::::::m  m:::::::mm           ss::::::::::s  h::::hh:::::hhh    
n::::::::::::::nn  e::::::eeeee:::::eem::::::::::mm::::::::::m        ss:::::::::::::s h::::::::::::::hh  
nn:::::::::::::::ne::::::e     e:::::em::::::::::::::::::::::m        s::::::ssss:::::sh:::::::hhh::::::h 
  n:::::nnnn:::::ne:::::::eeeee::::::em:::::mmm::::::mmm:::::m         s:::::s  ssssss h::::::h   h::::::h
  n::::n    n::::ne:::::::::::::::::e m::::m   m::::m   m::::m           s::::::s      h:::::h     h:::::h
  n::::n    n::::ne::::::eeeeeeeeeee  m::::m   m::::m   m::::m              s::::::s   h:::::h     h:::::h
  n::::n    n::::ne:::::::e           m::::m   m::::m   m::::m        ssssss   s:::::s h:::::h     h:::::h
  n::::n    n::::ne::::::::e          m::::m   m::::m   m::::m        s:::::ssss::::::sh:::::h     h:::::h
  n::::n    n::::n e::::::::eeeeeeee  m::::m   m::::m   m::::m ...... s::::::::::::::s h:::::h     h:::::h
  n::::n    n::::n  ee:::::::::::::e  m::::m   m::::m   m::::m .::::.  s:::::::::::ss  h:::::h     h:::::h
  nnnnnn    nnnnnn    eeeeeeeeeeeeee  mmmmmm   mmmmmm   mmmmmm ......   sssssssssss    hhhhhhh     hhhhhhh
	`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var dhash map[string]interface{}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nem.sh.json)")
	viper.SetDefault("node", "san.nem.ninja")
	viper.SetDefault("port", "7890")
	viper.SetDefault("protocol", "http")
	viper.SetDefault("address", "NCNUJTJ7HAYF6PV3ZPJIAYVJFTR4VT4FG4C4FRA5")
	viper.SetDefault("alias", dhash)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".nem.sh")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		viper.ConfigFileUsed()
	}
}
