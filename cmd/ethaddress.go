// Copyright © 2021 NAME HERE <EMAIL ADDRESS>
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
	"strings"

	"github.com/spf13/cobra"
	"github.com/tharsis/token/app"
)

// ethaddressCmd represents the ethaddress command
var ethaddressCmd = &cobra.Command{
	Use:   "ethaddress",
	Short: "ethaddress <mnemonic_words>",
	Long: `This command derives an Ethereum address from mnemonic words. It is done by converting mnemonic words into a private key that is used to derive Ethereum address.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mnemonic := strings.Join(args, " ")
		if err := app.DeriveEthAddressFrom(mnemonic); err != nil {
			return err
		}

		return nil
	},
	Args: cobra.ExactArgs(24),
}

func init() {
	RootCmd.AddCommand(ethaddressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ethaddressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ethaddressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
