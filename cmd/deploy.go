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
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tharsis/token/app"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy <mnemonic words>",
	Short: "deploy deploys a contract at Ethermint node",
	Long: `deploy uses mnemonic words to derive a private key. This private key is used to yield a public key and Ethereum address. 
	This Ethereum account acts as a fromAccount and pays gas costs to deploy a contract.`,
	RunE: runDeployCmd,
	Args: cobra.ExactArgs(24),
}

func runDeployCmd(_ *cobra.Command, args []string) error {
	mnemonic := strings.Join(args, " ")
	client, err := app.NewClient(mnemonic)
	if err != nil {
		return err
	}

	if err = client.DeployContract(); err != nil {
		return err
	}

	fmt.Println("deploy called")
	return nil
}

func init() {
	RootCmd.AddCommand(deployCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
