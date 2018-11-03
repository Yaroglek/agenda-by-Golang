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
	"agenda/service"
	"github.com/spf13/cobra"
)

// lCmd represents the l command
var lCmd = &cobra.Command{
	Use:   "l",
	Short: "login",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Login called")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		if username == "" || password == "" {
			fmt.Println("Please input username and password")
			return
		}
		if _, flag := service.GetCurUser(); flag == true {
			fmt.Println("Error: please logout")
			return
		}
		if tf := service.UserLogin(username, password); tf == true {
			fmt.Println("Login Successfully. Current User: ", username)
		} else {
			fmt.Println("Login fail: Wrong username or password")
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(lCmd)
	lCmd.Flags().StringP("username", "u", "", "agenda username")
	lCmd.Flags().StringP("password", "p","","agenda password")
}
