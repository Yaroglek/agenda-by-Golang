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

// duCmd represents the du command
var duCmd = &cobra.Command{
	Use:   "du",
	Short: "delete a user",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Delete User called")
		if user,flag := service.GetCurUser(); flag != true {
			fmt.Println("Error: please login")
		} else {
			if dflag := service.DeleteUser(user.Name); dflag != true {
				fmt.Println("Error! Please check error.log")
			} else {
				fmt.Println("Delete Sucessfully")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(duCmd)
}