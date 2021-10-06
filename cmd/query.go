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

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query <account_address> ",
	Short: "This command queries a token balance of either contract or account address",
	Long: `Firstly, it sets up a contract instance using a contract address provided. If it fails, it yield an error.
	 Otherwise, the command checks the number of arguments. If it has one argument, the command queries contract token supply. 
	 If it has 2 arguments, the command queries the balance of account address.`,
	RunE: runQueryCmd,
	Args: cobra.ExactArgs(1),
}

func runQueryCmd(cmd *cobra.Command, args []string) error {
	clientCtx, err := app.GetClientContext(cmd)
	if err != nil {
		return err
	}

	account := args[0]
	if err := clientCtx.Client.Query(account); err != nil {
		return err
	}

	return nil

	// if len(args) == 1 we display total supply of a contact
	/*
		if len(args) == 1 {
			totalSupply, err := instance.TotalSupply1(&bind.CallOpts{})
			if err != nil {
				return fmt.Errorf("TotalSupply1 err: %q", err)
			}

			fmt.Println("totalSupply of a contract", totalSupply)
			return nil
		}
	*/

}

func init() {
	RootCmd.AddCommand(queryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
