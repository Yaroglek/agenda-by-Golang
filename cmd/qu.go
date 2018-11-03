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

// fuCmd represents the fu command
var fuCmd = &cobra.Command{
	Use:   "qu",
	Short: "query users",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Find User called")
		ru := service.ListAllUser()
		for _, u := range ru {
			fmt.Println("----------------")
			fmt.Println("Username:", u.Name)
			fmt.Println("Phone:", u.Phone)
			fmt.Println("Email:", u.Email)
			fmt.Println("----------------")
		}
	},
}

func init() {
	rootCmd.AddCommand(fuCmd)
}