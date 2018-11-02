// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

	"github.com/spf13/cobra"
)

// rCmd represents the r command
var rCmd = &cobra.Command{
	Use:   "r",
	Short: "register",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")
		if username == "" {
			fmt.Println("Username is empty.")
		}
		if password == "" {
			fmt.Println("Password is empty.")
		}
		if email == "" {
			fmt.Println("Email is empty.")
		}
		if phone == "" {
			fmt.Println("Phone is empty.")
		}
	},
}

func init() {
	rootCmd.AddCommand(rCmd)
	rCmd.Flags().StringP("username", "u", "", "username")
	rCmd.Flags().StringP("password", "p", "", "password")
	rCmd.Flags().StringP("email", "e", "", "email")
	rCmd.Flags().StringP("phone", "p", "", "phone")
}
