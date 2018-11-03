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

// qmCmd represents the qm command
var qmCmd = &cobra.Command{
	Use:   "qm",
	Short: "quit a meeting",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Quit Meeting called")
		title, _ := cmd.Flags().GetString("title")
		if title == "" {
			fmt.Println("Please input meeting title")
			return
		}
		if user,flag := service.GetCurUser(); flag != true {
			fmt.Println("Error: please login")
		} else {
			if flag := service.QuitMeeting(user.Name, title); flag == false {
				fmt.Println("Error: Meeting not exist or you're not a participator of it")
			} else {
				fmt.Println("Quit Successfully")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(qmCmd)
	qmCmd.Flags().StringP("title", "t", "", "the title of meeting")
}
