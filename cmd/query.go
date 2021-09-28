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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/cobra"
	"github.com/tharsis/token/app"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
     token_address, address1
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: runQueryCmd,
	Args: cobra.RangeArgs(1, 2),
}

func runQueryCmd(cmd *cobra.Command, args []string) error {

	clientCtx, err := app.GetClientContext(cmd)
	if err != nil {
		return err
	}

	contractHexStr := args[0]
	client := clientCtx.Client
	instance, err := client.GetContractInstance(contractHexStr)
	if err != nil {
		return fmt.Errorf("GetContractInstance err: %q", err)
	}

	// if len(args) == 1 we display total supply of a contact
	if len(args) == 1 {
		totalSupply, err := instance.TotalSupply1(&bind.CallOpts{})
		if err != nil {
			return fmt.Errorf("TotalSupply1 err: %q", err)
		}

		fmt.Println("totalSupply of a contract", totalSupply)
		return nil
	}

	addrStr := args[1]
	// if len(args) == 2 , we display a total supply of an address args[1]
	bal, err := app.DisplayTokenBalance(instance, addrStr)
	if err != nil {
		return fmt.Errorf("DisplayTokenBalance err: %q", err)
	}

	fmt.Printf("Token balance of %s is: %d", args[1], bal)
	return nil
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
