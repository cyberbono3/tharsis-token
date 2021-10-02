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
	"github.com/spf13/cobra"

	"github.com/tharsis/token/app"
)

// mintCmd represents the mint command
var mintCmd = &cobra.Command{
	Use:   "mint <ethereum_address> <amount> ",
	Short: "Mint <amount> of tokens in <ethereum address> ",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: runMintCmd,
	Args: cobra.ExactArgs(2),
}

func runMintCmd(cmd *cobra.Command, args []string) error {

	clientCtx, err := app.GetClientContext(cmd)
	if err != nil {
		return err
	}

	to := args[0]
	amountStr := args[1]

	if err = clientCtx.Client.Mint(to, amountStr); err != nil {
		return err
	}

	return nil
}

func init() {
	RootCmd.AddCommand(mintCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mintCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mintCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
