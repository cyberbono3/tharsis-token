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
	"math/big"

	"github.com/spf13/cobra"
	"github.com/tharsis/token/app"
)

// transferCmd represents the transfer command
var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer tokens between two accounts ",
	Long: `Transfer tokens between two accounts that can be minted using deployed ERC-20 contract`,
	RunE: runTransferCmd,
	Args: cobra.ExactArgs(2),
}

func runTransferCmd(cmd *cobra.Command, args []string) error {
	clientCtx, err := app.GetClientContext(cmd) 
	if err != nil {
		return err 
	}

	client := clientCtx.Client
	to := args[0]
	tokens := args[1]

	// convert string to *big.Int
	n := new(big.Int)
    n, ok := n.SetString(tokens, 10)
    if !ok {
        fmt.Println("SetString: error")
        return err
    }
   
	if err = client.TransferTokens(to, n); err != nil {
		return err
	}

	fmt.Println("transfer has been executed successfuly")

	return nil
}

func init() {
	RootCmd.AddCommand(transferCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transferCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transferCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
